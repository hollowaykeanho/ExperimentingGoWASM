# Holloway's GoWASM Experiments
![banner](artwork/banner-1200x628.webp)

This is the principle repository for (Holloway) Chew, Kean Ho's experimentation
to develop a pure Go web application supporting
[Progressive Web Application](https://web.dev/progressive-web-apps/) offline
mode alongside using pure Go to develop web application entirely.




## Experiemntation Goals
The goal of unifying web application development for client-side rendering using
a proper programming language like [Go](https://go.dev/) has been aspired but
low implementations due to little efforts. That's why the primary goal is to:

1. Develop the foundational Go library so that the customers (Go web developer)
   to deploy web app.
2. Develop client-side rendering capability using [Hugo](https://gohugo.io/) for
   cost reduction between the server sides and the end-user sides.
3. Develop the necessary CSS/Sass frontend rendering libraries to keep the
   foundation reasonably and visually appealing at minimum.
4. Use scalable tools that can roll out updates for multiple technologies
   modularly and without much fears.




## Local and Offline Testing
The repository is designed in a way that it can operate in offline or
bad connectivity environments. To do that, simple have <code>hugo</code>
(https://gohugo.io/) installed in your system and then execute:

```
.configs/hugo/server.cmd
```

Once done, please visit the URL site presented in the terminal. The default
is: http://localhost:8080




## License
The repository is licensed under MIT License.
