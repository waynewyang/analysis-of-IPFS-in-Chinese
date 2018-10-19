# DNS

## 目录
- [基础知识](#基础知识)
- [目标](#目标)
- [代码分析](#代码分析)
- [操作示例](#操作示例)

## 基础知识
- 常用域名记录解释：A记录、MX记录、CNAME记录、TXT记录、AAAA记录、NS记录
- A记录
	A记录是用来创建到IPV4地址的记录。
	在命令行下可以通过nslookup -qt=a www.ezloo.com 来查看A记录。

- AAAA记录
	AAAA记录是一个指向IPv6地址的记录。

- MX记录
	在命令行下可以通过 nslookup -qt=mx ezloo.com 来查看MX记录。
	mx 记录的权重对 Mail 服务是很重要的，当发送邮件时，Mail 服务器先对域名进行解析，查找 mx 记录。先找权重数最小的服务器（比如说是 10），如果能连通，那么就将服务器发送过去；如果无法连通 mx 记录为 10 的服务器，那么才将邮件发送到权重为 20 的 mail 服务器上。

- CNAME记录
	CNAME记录也成别名记录，它允许你将多个记录映射到同一台计算机上。比如你建了如下几条记录：
	a1 CNAME a.ezloo.com 
	a2 CNAME a.ezloo.com 
	a3 CNAME a.ezloo.com 
	a A 111.222.111.222
	我们访问a1（a2，a3）.ezloo.com的时候，域名解析服务器会返回一个CNAME记录，并且指向a.ezloo.com，然后我们的本地电脑会再发送一个请求，请求a.ezloo.com的解析，返回IP地址。
	当我们要指向很多的域名到一台电脑上的时候，用CNAME比较方便，就如上面的例子，我们如果服务器更换IP了，我们只要更换a.ezloo.com的A记录即可。
	在命令行下可以使用nslookup -qt=cname a.ezloo.com来查看CNAME记录。

- TXT记录
	TXT记录一般是为某条记录设置说明，比如你新建了一条a.ezloo.com的TXT记录，TXT记录内容"this is a test TXT record."，然后你用 nslookup -qt=txt a.ezloo.com ，你就能看到"this is a test TXT record"的字样。
	在命令行下可以使用nslookup -qt=txt a.ezloo.com来查看TXT记录。

- NS记录
NS记录是域名服务器记录，用来指定域名由哪台服务器来进行解析。可以使用nslookup -qt=ns ezloo.com来查看。

- [回到目录](#目录)

## 目标
- 基本原理
	利用TXT记录，将通用人类可识别域名，解析未IPFS or IPNS地址
	添加TXT记录的时候主机记录为   _dnslink.你的域名 或者你的域名
	
- 域名解析为ipfs地址，示例
```
			"ipfs.example.com": []string{
				"dnslink=/ipfs/QmY3hE8xgFCjGcz6PHgnvJz5HZi1BaKRfPkn1ghZUcYMjD",
			},
			"_dnslink.dipfs.example.com": []string{
				"dnslink=/ipfs/QmY3hE8xgFCjGcz6PHgnvJz5HZi1BaKRfPkn1ghZUcYMjD",
			},
			"dns1.example.com": []string{
				"dnslink=/ipns/ipfs.example.com",
			},
			"dns2.example.com": []string{
				"dnslink=/ipns/dns1.example.com",
			},
```

- [回到目录](#目录)

## 代码分析
- base.go
```
//解析接口
type resolver interface {
	// resolveOnce looks up a name once (without recursion).
	resolveOnce(ctx context.Context, name string, options *opts.ResolveOpts) (value path.Path, ttl time.Duration, err error)
}

// 解析ipns的上层方法,name 表示的是通用域名
// resolve is a helper for implementing Resolver.ResolveN using resolveOnce.
func resolve(ctx context.Context, r resolver, name string, options *opts.ResolveOpts, prefixes ...string) (path.Path, error) 
```
- dns.go
```
package namesys

import (
	"context"
	"errors"
	"net"
	"strings"
	"time"

	opts "github.com/ipfs/go-ipfs/namesys/opts"
	path "gx/ipfs/QmTKaiDxQqVxmA1bRipSuP7hnTSgnMSmEa98NYeS6fcoiv/go-path"
	isd "gx/ipfs/QmZmmuAXgX73UQmX1jRKjTGmjzq24Jinqkq8vzkBtno4uX/go-is-domain"
)

type LookupTXTFunc func(name string) (txt []string, err error)

// DNSResolver implements a Resolver on DNS domains
type DNSResolver struct {
	// http的TXT域名解析,返回格式应该为ipfs(ipns)或者dnslink=/ipfs(ipns)
	lookupTXT LookupTXTFunc
	// TODO: maybe some sort of caching?
	// cache would need a timeout
}

// NewDNSResolver constructs a name resolver using DNS TXT records.
func NewDNSResolver() *DNSResolver {
	return &DNSResolver{lookupTXT: net.LookupTXT}
}

// Resolve implements Resolver.
func (r *DNSResolver) Resolve(ctx context.Context, name string, options ...opts.ResolveOpt) (path.Path, error) {
	//会调用resolveOnce
	return resolve(ctx, r, name, opts.ProcessOpts(options), "/ipns/")
}

// ipfs or ipns地址
type lookupRes struct {
	path  path.Path
	error error
}

// resolveOnce implements resolver.
// TXT records for a given domain name should contain a b58
// encoded multihash.
func (r *DNSResolver) resolveOnce(ctx context.Context, name string, options *opts.ResolveOpts) (path.Path, time.Duration, error) {
	// 去掉后缀
	segments := strings.SplitN(name, "/", 2)
	domain := segments[0]

	if !isd.IsDomain(domain) {
		return "", 0, errors.New("not a valid domain name")
	}
	log.Debugf("DNSResolver resolving %s", domain)

	rootChan := make(chan lookupRes, 1)
	go workDomain(r, domain, rootChan)

	subChan := make(chan lookupRes, 1)
	go workDomain(r, "_dnslink."+domain, subChan)

	var subRes lookupRes
	select {
	case subRes = <-subChan:
	case <-ctx.Done():
		return "", 0, ctx.Err()
	}

	var p path.Path
	if subRes.error == nil {
		p = subRes.path
	} else {
		var rootRes lookupRes
		select {
		case rootRes = <-rootChan:
		case <-ctx.Done():
			return "", 0, ctx.Err()
		}
		if rootRes.error == nil {
			p = rootRes.path
		} else {
			return "", 0, ErrResolveFailed
		}
	}
	var err error
	if len(segments) > 1 {
		p, err = path.FromSegments("", strings.TrimRight(p.String(), "/"), segments[1])
	}
	return p, 0, err
}

func workDomain(r *DNSResolver, name string, res chan lookupRes) {
	//http url解析
	txt, err := r.lookupTXT(name)

	if err != nil {
		// Error is != nil
		res <- lookupRes{"", err}
		return
	}

	for _, t := range txt {
		p, err := parseEntry(t)
		if err == nil {
			res <- lookupRes{p, nil}
			return
		}
	}
	res <- lookupRes{"", ErrResolveFailed}
}

func parseEntry(txt string) (path.Path, error) {
	p, err := path.ParseCidToPath(txt) // bare IPFS multihashes
	if err == nil {                    //如果地址txt为cid的最终地址格式,直接返回
		return p, nil
	}

	//如果是dnslink（表示还没有解析完成，需要继续解析）
	return tryParseDnsLink(txt)
}

func tryParseDnsLink(txt string) (path.Path, error) {
	parts := strings.SplitN(txt, "=", 2)
	if len(parts) == 2 && parts[0] == "dnslink" {
		return path.ParsePath(parts[1])
	}

	return "", errors.New("not a valid dnslink entry")
}
```

- [回到目录](#目录)


## 操作示例

- 待添加
- [回到目录](#目录)