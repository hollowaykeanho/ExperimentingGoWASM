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
   1. **[ DONE ]** - Power on with `ExperimentingGoWASM` Go workspace with
      `vendor/hestiaGo` and `vendor/presentoGo` packages.
   2. **[ DONE ]** - Develop Monteur's test, build, package, and release
      recipes for GoWASM.
   3. Develop baseline foundation for using GoWASM.
   4. **[ DONE ]** - Develop `<body>` manipulations directly in GoWASM
      (output).
   5. **[ DONE ]** - Develop `<body>` manipulations directly in GoWASM (input).
   6. **[ DONE ]** - Develop `<body>` manipulations directly in GoWASM
      (event-driven).
   7. **[ DONE ]** Explore and conclude the necessity of using
      [TinyGo](https://tinygo.org/).
2. Develop client-side rendering capability using [Hugo](https://gohugo.io/) for
   cost reduction between the server sides and the end-user sides.
   1. **[ DONE ]** - Power on with Presento Theme with NoCSS but rendering
      capabilities.
   2. **[ DONE ]** - Develop necessary partial functions to isolate Hugo's risky
      functions and also to prevent supply-chain vendor locked-in whenever
      possible.
   3. **[ DONE ]** - Develop plain HTML+Javascript pages necessary to bring up
      Go-WASM.
   4. **[ DONE ]** - Power on Go-WASM+HTML+Javascript on Hugo.
   5. **[ DONE ]** - Deploy Gunzip against Go-WASM to reduce its size with
      Monteur.
   6. **[ DOING ]** - Develop necessary foundation for Go-WASM client-side rendering.
3. Develop the necessary CSS/Sass frontend rendering libraries to keep the
   foundation reasonably and visually appealing at minimum.
   1. **[ DONE ]** - Analyize (1) and (2) to see exactly where to develop the
      CSS/Sass and why. Make sure it's documentable.
   2. **[ DOING ]** - Develop the core component to render the page without device
      screen locking.
4. Use scalable tools that can roll out updates for multiple technologies
   modularly and without much fears.
   1. **[ DONE ]** - deployed ZORALab's Monteur to manage the repository
      systematically, customizable at scale.
   2. **[ DONE ]** - deployed Hugo Themes module to manage Hugo specific and
      scalable setup.




## 3rd-Party Dependencies
Here are the list of 3rd-party dependencies used so far:

1. ZORALab's Monteur - https://monteur.zoralab.com
2. Go - https://go.dev
3. Hugo - https://gohugo.io/
4. GolangCI-Lint - https://golangci-lint.run/
5. TinyGo - https://tinygo.org/docs/guides/webassembly/



### About GZip Output Files
After decent research, it was found that many network CDN and load balancer
service providers (e.g. CloudFlare) automatically serves the content in
compressed mode with GZip or Brotil (see
https://support.cloudflare.com/hc/en-us/articles/200168396-Does-Cloudflare-compress-resources-).

For self-hosting CDN servers like Nginx or Apache, they already are serving
compression executions when enabled
(see: https://docs.nginx.com/nginx/admin-guide/web-server/compression/).

Hence, there is no need to worry about manual compression.




## Deployment Requirements
To determine where pure-Go WASM rendering to serve, a detailed scan through
of how consumption behavior affects the web technologies and how it serves.


### Web Format Availability
Here, we look into what and how consumptions influences the web languages:

| Available Formats   | Caused By    | Content By | Designed By | Behavior By |
|:--------------------|:------------:|:----------:|:-----------:|:-----------:|
| `html`, `css`, `js` | normal use   | `html`     | `css`       | `js`        |
| `html`, `css`       | js-block     | `html`     | `css`       | `css`       |
| `html`, `css`       | ad-block     | `html`     | `css`       | `css`       |
| `html`, `js`        | React-type   | `html`     | `js`        | `js`        |
| `html`, `js`        | Angular-type | `html`     | `js`        | `js`        |
|  `js`               | single-page  | `js`       | `js`        | `js`        |
|  `html`             | baremetal    | `html`     |             |             |

Conclusion for Designing Pure-Go WASM Library
1. **The UI library must be capable of generating its necessary HTML content codes**.
2. **The UI library must be capable of generating its necessary CSS styling codes**.
   1. **The CSS codes MUST be flexible for various deployments** notably as:
       1. inline HTML
       2. as `css` file (external)
       3. inside `js`
       4. as inline `css` (inside HTML)
       5. `js` rendering codes.
3. **The content codes must be flexible enough to store in both `html` and `js` formats**.
4. the behavior part is split into 2 parts:
   1. **UI related behavior (e.g. bouncing, dancing, etc) shall be handled solely
      by `css`**.
   2. **Data processing or routing (e.g. trigger loading screen, transmitting
      submission to server, etc) should be solely handled by `js`**.
5. **The UI library MUST leverage the compiler's unused code shedding capability** for
   any of the user choice of output.

### UI Tech Availability
Here, we look into different UI tech availability for rendering the a display output:

| Technologies  | Source                                     |
|:--------------|:-------------------------------------------|
| `vanilla web` | https://html.spec.whatwg.org/#introduction |
| `qt`          | https://doc.qt.io/qt-6/wasm.html           |
| `wails`       | https://wails.io/                          |
| `lorca`       | https://github.com/zserge/lorca            |
| `pwa`         | https://web.dev/progressive-web-apps/      |
| `webview`     | https://github.com/webview/webview         |
| `seed`        | https://github.com/qlova/seed              |
| `qt` (Go)     | https://github.com/therecipe/qt            |
| `tcell`       | https://github.com/gdamore/tcell           |

Conclusion:

1. The UI package **MUST leave room for horizontal scaling outside of
   web technologies alone**.
2. The package **MUST be a factory of similar design to generate the
   technology specific codes while allowing the compiler to shed unused
   ones seamlessly**.

### Rendering Devices
Here, we look into different devices rendering the same content page:


| Technologies                     | Type      | Display Size (width x height) | Pixel Density | Source         |
|:---------------------------------|:----------|:-----------------------------:|:-------------:|:--------------:|
| Apple Watch 6 (40mm)             | Wearables | `162px x 197px`               | 2             | https://www.webmobilefirst.com/en/devices/apple-watch-serie-6/ |
| Apple iPhone SE (2018)           | Mobile    | `320px x 568px`               | 2             | https://www.webmobilefirst.com/en/devices/apple-iphone-se/ |
| Apple iPhone 5s                  | Mobile    | `320px x 568px`               | 2             | https://www.webmobilefirst.com/en/devices/apple-iphone-5s/ |
| Sony Experia XZ                  | Mobile    | `320px x 640px`               | 3             | https://www.webmobilefirst.com/en/devices/sony-xperia-xz/ |
| Huawei Mate 10 Lite              | Mobile    | `360px x 720px`               | 3             | https://www.webmobilefirst.com/en/devices/huawei-mate-10-lite/ |
| Samsung Galaxy S22 Ultra         | Mobile    | `360px x 772px`               | 4             | https://www.webmobilefirst.com/en/devices/samsung-galaxy-s22-ultra-2022/ |
| Samsung Galaxy Ultra 22+         | Mobile    | `360px x 780px`               | 3             | https://www.webmobilefirst.com/en/devices/samsung-galaxy-s22-plus-2022/ |
| Honor 9x                         | Mobile    | `360px x 780px`               | 3             | https://www.webmobilefirst.com/en/devices/honor-9x/ |
| Huawei Mate 30                   | Mobile    | `360px x 780px`               | 3             | https://www.webmobilefirst.com/en/devices/huawei-mate-30/ |
| Huawei P30 Pro                   | Mobile    | `360px x 780px`               | 3             | https://www.webmobilefirst.com/en/devices/huawei-p30-pro/ |
| Google Pixel 6 Pro               | Mobile    | `360px x 780px`               | 4             | https://www.webmobilefirst.com/en/devices/google-pixel-6-pro/ |
| Xiaomi 12                        | Mobile    | `360px x 800px`               | 3             | https://www.webmobilefirst.com/en/devices/xiaomi-12-2022/ |
| Xiaomi 11i                       | Mobile    | `360px x 800px`               | 3             | https://www.webmobilefirst.com/en/devices/xiaomi-mi-11i/ |
| Generic Android Phone            | Mobile    | `360px x 800px`               | 2             | https://www.webmobilefirst.com/en/devices/non-branded-android-smartphone/ |
| Apple iPhone 7                   | Mobile    | `375px x 667px`               | 2             | https://www.webmobilefirst.com/en/devices/apple-iphone-7/ |
| iPhone 13 Mini (2021)            | Mobile    | `375px x 812px`               | 3             | https://www.webmobilefirst.com/en/devices/apple-iphone-13-mini-2021/ |
| Huawei Mate 30 Pro               | Mobile    | `392px x 800px`               | 3             | https://www.webmobilefirst.com/en/devices/huawei-mate-30-pro/ |
| Apple iPhone 7 Plus              | Mobile    | `414px x 736px`               | 3             | https://www.webmobilefirst.com/en/devices/apple-iphone-7-plus/ |
| Samsung Galaxy Note20 Ultra      | Mobile    | `414px x 883px`               | 3.5           | https://www.webmobilefirst.com/en/devices/samsung-galaxy-note20-ultra/ |
| Apple iPhone XS Max              | Mobile    | `414px x 896px`               | 3             | https://www.webmobilefirst.com/en/devices/apple-iphone-xs-max/ |
| Samsung Galaxy Fold (2019)       | Mobile    | `768px x 1076px`              | 2             | https://www.webmobilefirst.com/en/devices/samsung-galaxy-fold/ |
| Samsung Galaxy Fold (2020)       | Mobile    | `884px x 1104px`              | 2             | https://www.webmobilefirst.com/en/devices/samsung-galaxy-fold2/ |
| Amazon Fire 7 (2017)             | Tablet    | `1024px x 600px`              | 1             | http://responsivechecker.net/device/fire-7-2017 |
| Samsung Galaxy Tab S3 9.7        | Tablet    | `1024px x 768px`              | 2             | http://responsivechecker.net/device/galaxy-tab-s3-9-7 |
| Apple iPad Mini 4                | Tablet    | `1024px x 768px`              | 2             | http://responsivechecker.net/device/ipad-mini-4 |
| Microsoft Surface Duo            | Mobile    | `1114px x 705px`              | 2.5           | https://www.webmobilefirst.com/en/devices/microsoft-surface-duo/ |
| Apple iPad 4                     | Tablet    | `1180px x 820px`              | 2             | https://www.webmobilefirst.com/en/devices/apple-ipad-air-4/ |
| Amazon Fire HD 10 (2017)         | Tablet    | `1280px x 800px`              | 2             | http://responsivechecker.net/device/fire-hd-10-2017 |
| Macbook Air 2020 13"             | Laptop    | `1280px x 800px`              | 2             | https://www.webmobilefirst.com/en/devices/macbook-air/ |
| iPad Pro                         | Tablet    | `1366px x 1024px`             | 2             | http://responsivechecker.net/device/ipad-pro/ |
| iPad Pro 12.9"                   | Tablet    | `1366px x 1024px`             | 2             | http://responsivechecker.net/device/ipad-pro-12-9 |
| Macbook Pro                      | Laptop    | `1728px x 1117px`             | 2             | https://www.webmobilefirst.com/en/devices/apple-macbook-pro-16-2021/ |
| Samsung Smart TV NEO QLED 4k 55" | TV        | `1920px x 1080px`             | 2             | https://www.webmobilefirst.com/en/devices/samsung-smart-tv-neo-qled-4k-55/ |
| iMac 24" 2021                    | Desktop   | `2048px x 1152px`             | 2             | https://www.webmobilefirst.com/en/devices/apple-imac-24-inch-2021/ |

Conclusion:

1. **UI library web rendering SHALL NOT DEPENDS on device-oriented
   media query breakpoint**.
2. Any media query **SHALL be oriented to the design module with its
   own definitions of breakpoint indepnendent of device dimension**.
3. As the hardware industry rolls out odd screen sizes, **there is a
   need to develop a new approach to avoid breakpoint dimension whenever
   possible**.
4. **All media tags (e.g. images, videos, etc) shall have pixel density
   awareness**.




## Critical Issues
Here are the tracking issues critical to the research success:

1. Tinygo WASM Memory Leak (first detected after implementing Chain kernel
   example) - https://github.com/tinygo-org/tinygo/issues/1140




## (1) Utilize Local CI - Monteur
The FIRST technology chosen was [ZORALab's Monteur](https://monteur.zoralab.com)
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
The SECOND technology selected was `hugo` for simple static site generations.
The repository is designed in a way that it can operate in offline or bad
connectivity environments. To do that, execute:

```
.configs/hugo/server.cmd
```

Once done, please visit the URL site presented in the terminal. The default
is: http://localhost:8080

If you need to work on something else, you need to setup a new terminal again.




## (4) GoWASM Go Workspace
The THIRD technologies selected were `go` and `golangci-lint` for their
simplicity and portability sake. In order to work on the Go project, you need
to open a new terminal, perform the Step (2) for it, and then change directory
into `wasmExpGo`. That is the root location of the Go source codes and its
workspace.

```
cd wasmExpGo
```

While the workspace is for this repository, the file structures are arranged
into 3 different Go `modules`: `wasmExpGo`, `hestiaGo`, and `presentoGo`. The
minimal file structure is as follows:

```
.
├── app
│   └── wasm
│       └── main.go
├── go.mod
├── hestiaGo
│   ├── go.mod
│   └── version.go
└── presentoGo
    ├── go.mod
    └── version.go
```

Each module has its specific roles to prevent vendor locked-in threat while
maintaining Go's quality in package modularlities. They key roles are:

1. `wasmExpGo` - the repository specific packages.
2. `hestiaGo` - the common library packages that will be upstreamed to:
   https://github.com/ZORALab/Hestia
3. `presentoGo` - the frontend UI rendering packages that will be upstreamed to:
   https://github.com/ZORALab/Presento

To prevent supply-chain nightmare from affecting any project, `go.mod` plays a
critical role here to make sure all the source codes know where to find the
packages.

Generally speaking, all import statements **SHALL use the `replace` clause**
to denote its sourcing navigations. Here is a simple example:

```
// Copyright 2022 "Holloway" Chew, Kean Ho <hollowaykeanho@gmail.com>
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to
// deal in the Software without restriction, including without limitation the
// rights to use, copy, modify, merge, publish, distribute, sublicense, and/or
// sell copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING
// FROM, OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER
// DEALINGS IN THE SOFTWARE.

package main

import (
	"fmt"
	"hestiaGo"
	"presentoGo"
)

func main() {
	fmt.Printf("Hello World\n")
	fmt.Printf("Hestia Version: %s\n", hestiaGo.VERSION)
	fmt.Printf("Presento Version: %s\n", presentoGo.VERSION)
}
```

Notice that both `hestiaGo` and `presentoGo` are not imported using
`github.com/hollowaykeanho/...`. Instead, they are using their navigator.

Then in `wasmExpGo`'s `go.mod`, it has replace clauses as follows:

```
module github.com/hollowaykeanho/ExperimentingGoWASM/wasmExpGo

go 1.18

replace (
	github.com/hollowaykeanho/ExperimentingGoWASM/wasmExpGo => ./
	hestiaGo => ./hestiaGo
	presentoGo => ./presentoGo
	wasmExpGo => ./
)
```

The `module` clause remains as default since we want to maintain `go get`
capability. However, in source codes, we use `wasmExpGo` instead just like the
other 2.

Then, in each 3rd-party libraries, their modules are in accordance to their
respective `go.mod` settings:

```
module https://github.com/ZORALab/Hestia/hestiaGo

go 1.18

replace (
	github.com/ZORALab/Hestia/hestiaGo => ./
	hestiaGo => ./
)
```




## (5) Testing Go Workspace
With Monteur Test CI Job made available, given the correct recipe, Monteur
can perform either pinpoint or recursive testing against Go packges. All the
user needs to do is:

```
$ monteur test
```

And then all the necessary result data files are generated into the package
directory itself.




## (6) Building Go WASM
With Monteur Build CI Job made available, given the correct recipe, Monteur
can peform percise Go build in a reproducible manner. All the user needs to do
is:

```
$ monteur build
```




## (7) Package Go WASM
With Monteur Package CI Job made available, given the correct recipe, Monteur
can peform proper packaging in a reproducible manner. All the user needs to do
is:

```
$ monteur package
```




## (8) Release to Hugo
With Monteur Release CI Job made available, given the correct recipe, Monteur
can peform proper manage the Go WASM systematically. All the user needs to do
is:

```
$ monteur release
```

At this point, Go development is considered completed.




## (9) Compose Static Site Generations
With Monteur Compose CI Job made available, given the correct recipe, Monteur
can properly compose Hugo's website artifact for web publications. All the user
needs to do is:

```
$ monteur compose
```

At this point, Hugo development is considered completed.




## (10) Publish the Static Site Artifact
With Monteur Publish CI Job made available, given the correct recipe, Monteur
can properly publish the composed website artifact to the public. All the user
needs to do is:

```
$ monteur publish
```

At this point, the development is considered deployed.




## License
The repository is licensed under MIT License.

