Recently started thinking about next.js and it's pitfalls. eg background/long tasks like notifications etc.

Came across qstash which basically just echo's the request you send it to another endpoint.
Qstash has much more funtionality like scheduling/delays/batching etc but fundamentally it's just an echoer

With the nextjs issue, we can use qstash to send background jobs to a nextjs api route which can do the work there
and the user won't be blocked and have a fast experience. Still nextjs api routes are pretty limited on a free plan, when hosted with vercel,
and even then I think it maxes out at 900s on an enterprise plan 

Might try creating a qstash clone in golang just for fun and learning.

I've recently been thinking about AI intergrations
and sometimes they may take a while, especially if you want to render images/video. Could be a fun problem to solve.

Just started playing around with how golang would do background jobs, creating a pool of go routines so that jobs/tasks
can be done in parallel instead of sequentially. If some job was to take minutes then the rest of them would be blocked.
Also learned about channel capacity, if I tried to enqueue more than 100 jobs then the subsequent requests would just wait
until they timeout (if they timed out). This allows the queue to have some support and not get overwhelmed. In my case it was just
an in-memory queue but for production might use a proper messaging/queuing technology like kafka/redis/rabbitMQ. 

Could add an abstraction layer interface which has a few functions and behind that I plug in whatever implementation I want.

Regardless, fun experiment and learning experience.

Takeaways
* For MVP/starting out just have all of the jobs in next api routes with qstash
* Migrate to a dedicated server if background jobs start taking a while
