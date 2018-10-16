# DNS

## 目标
- 域名解析为ipfs地址（dns服务器，未ready？）
    解析对象及结果如下所示
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

## 代码分析
- base.go
```
//解析接口
type resolver interface {
	// resolveOnce looks up a name once (without recursion).
	resolveOnce(ctx context.Context, name string, options *opts.ResolveOpts) (value path.Path, ttl time.Duration, err error)
}

// 解析ipns的上层方法,name 表示的是http域名
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
	// http的域名解析,返回格式应该为ipfs(ipns)或者dnslink=/ipfs(ipns)
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

	//如果是dnslink继续解析
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
