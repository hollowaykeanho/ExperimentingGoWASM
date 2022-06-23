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
   1. **[ DOING ]** - Power on with `ExperimentingGoWASM` Go workspace with
      `vendor/hestiaGo` and `vendor/presentoGo` packages.
   2. Develop Monteur's test, build, package, and release recipes for GoWASM.
   3. Develop baseline foundation for using GoWASM.
   4. Explore and conclude the necessity of using [TinyGo](https://tinygo.org/).
2. Develop client-side rendering capability using [Hugo](https://gohugo.io/) for
   cost reduction between the server sides and the end-user sides.
   1. **[ DONE ]** - Power on with Presento Theme with NoCSS but rendering
      capabilities.
   2. Develop necessary partial functions to isolate Hugo functions (prevent
      vendor locked-in).
   3. Develop plain HTML+Javascript pages necessary to bring up Go-WASM.
   4. Power on Go-WASM+HTML+Javascript on Hugo.
   5. Deploy Gunzip against Go-WASM to reduce its size with Monteur.
   6. Develop necessary foundation for Go-WASM client-side rendering.
3. Develop the necessary CSS/Sass frontend rendering libraries to keep the
   foundation reasonably and visually appealing at minimum.
   1. Analyize (1) and (2) to see exactly where to develop the CSS/Sass
      and why. Make sure it's documentable.
   2. Develop the core component to render the page without device screen
      locking.
4. Use scalable tools that can roll out updates for multiple technologies
   modularly and without much fears.
   1. **[ DONE ]** - deployed ZORALab's Monteur to manage the repository
      systematically, customizable at scale.
   2. **[ DONE ]** - deployed Hugo Themes module to manage Hugo specific and
      scalable setup.




## (1) Utilize Local CI - Monteur
The first technology chosen was [ZORALab's Monteur](https://monteur.zoralab.com)
to manage the repository development continuously and controls with confidences
when deploying the repository in a decentralized manner.

The first step you need to do would be installing Monteur as per instructed in
their official website.

Once the `monteur` program is available, proceed to your own copy and perform:

```bash
$ monteur setup
```

Monteur shall setup all the repository's dependencies and configurations
seamlessly. Repeat this `monteur setup` command whenever there is an update from
Monteur Setup Job recipes OR something went wrong with the current setup in
your repository.




## (2) Local Development
The second step is to bring up your local development. Monteur setup local
filesystem inside the repository. Hence, whenever you open a new terminal to
wanting to develop this repository, simply do the following:

```bash
$ source .monteurFS/config/main
```

Once done, the terminal you're in is now configured to be repository specific.
You're now ready to develop the repository. All instructions beyond this step
**assumes you always this step**.

> **NOTE**
>
> Whenever you find any software that is missing in action, you're likely
> forgotten this step.




## (3) Local Hosting Hugo
The repository is designed in a way that it can operate in offline or
bad connectivity environments. To do that, execute:

```
.configs/hugo/server.cmd
```

Once done, please visit the URL site presented in the terminal. The default
is: http://localhost:8080




## License
The repository is licensed under MIT License.
