## context

### 示例控制线程的方式
- WaitGroup等待
>好多个goroutine协同做一件事情的时候，因为每个goroutine做的都是这件事情的一部分，只有全部的goroutine都完成，这件事情才算是完成，这是等待的方式。

- Chan通知


- Context 灵活多线程多级控制
