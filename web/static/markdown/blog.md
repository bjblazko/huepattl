# Timo's blog

## Stardate 2024-02-03, Saturday: Captain's Log 

So after some twenty years of not really having a web site and being a hobbyist and
professional software developer... I think it is time again to actually run a site again.
I mean - isn't it a question of "pride"? ;)

So, as of now I do not know what will come here in the future. Maybe just a blog,
maybe I will provide information about my current and future projects here.

To be honest, my first idea to have a website again was not to have a website again,
but I wanted to develop one. From scratch, just for the fun of it and for learning.

Well, currently, I am just rendering markdown based text here - dynamically rendered
each time you visit this page. There are much smarter ways of achieving the goal...

I could have done this instead:

- build a site with Github Pages
- use some other hosted CMS such as Wordpress, Joomla...
- using a platform such as Squarespace, Wix, IONOS etc.
- use a professional static page generator such as Hugo
- just create some static HTML files and throw them on a VM in some datacenter
- spin up a Raspberry Pi with or without runnung Docker, doing dynamic DNS routing to that box and run the pages using e.g.
  - JVM with Quarkus, Spring Boot, Micronaut, Apache Tomcat, or just plain Jetty (JVM was and is my go-to stack for the past 22 years)
  - doing similar but natively, such as building my Quarkus app using GraalVM
  - spin up some Ruby, PHP, Django, LAMP, Sinatra, Node.js etc.
- doing the above in some datacenter or even cloud provider
- ...

Instead - honestly - on the one hand I was tired of the operations part. I did not want to
get my VM up to date, maintain and patch OS and middleware, play around with firewalls.
Tons of other stuff I am forgetting right now.

So, I decided to go to the dark side and host my site with a cloud provider. The serverless
way. So, no dedicated Kubernetes cluster or VM, just plain Google Cloud Run.
This is basically a Kubernetes Knative deployment, running on an anonymous Kubernetes Cluster
that Google operates and I am not aware of.

You develop something, make a Dockerfile, and deploy it with some terminal command (in a pipleine or not).

That's for the ops part, I think more on that later.

OK, so I decided to program my own stuff - knowing it would be inferior to anything else
but it is just fun. I decided to leave my feeling-at-home stack (JVM, Quarkus, Kotlin)
and use something different. Something I was aware of and being one of the few stack
I intentionally ignored: Go (golang).

Why? I mean there are other great programming languages, runtimes and frameworks - but
come on... Go uses capitals for functions! In Go you have `func DoSomething(arg string)` instead
of `fun doSomething(arg: String)` like in most other languages.

Verbs that start with uppercase letters! I had to ignore Go. Go uses upper and lower case
to define the visibility scope of something. While other languages use reserved words like
`public` and `private` for that, Go dooes this using case.

I could have accepted all kind of styles differing from Java or Kotlin such as writing in
Kebap case, snake case etc. I get along with doing OO or functional... but this?

Well, here we are, this page is presented employing Go. But that is up for future entries here.