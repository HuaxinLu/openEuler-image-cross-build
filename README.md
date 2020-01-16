# openEuler-image-cross-build
This is the second project of [**openEuler hackathon game**](https://openeuler.org/zh/events/2020hdc.html#section5). This project is a tool to build the cross-platform docker images. At the moment, you can use this project to build an aarch64 docker image in the x86 platform without any modification of Dockerfile.

## Prerequisites
This project is developed and tested in Huawei cloud server, and the OS is CentOS 7.6-x86_64. Docker daemon is required when the software is running. Docker 19.03.5 is recommended.

## How to use it
1. Run `sh prepare.sh` to get **qemu-aarch64-static** and set **binfmt_misc** function.
2. Copy **qemu-aarch64-static** to the directory that you want to build a docker image.
3. Run `go install` to install the project.
4. You can run `qemu-aarch64-static pull imagename`, `qemu-aarch64-static push imagename` and `qemu-aarch64-static build -f dockerfilename -t imagetag directory` like docker.

## Example
The **example** folder is an example to use this project to build an aarch64 httpd image based on openEuler 1.0 OS image in x86_64 platform.
1. Run docker daemon and finish prepare work as **How to use it** chapter.
2. Run `openEuler-image-cross-build pull huaxinlu/openeulergame2019:v3.1`. This is a public image, so the username and password are empty.
3. Copy **qemu-aarch64-static** to the **example** folder and run `openEuler-image-cross-build build -t httpd-openeuler-example example/` to build the image.
4. Because the **binfmt_misc** is set, you can test this image in host. Firstly, run the image using `docker run -d -p 80:80 httpd-openeuler-example`. Then run `curl http://localhost:80`, and "Hello World!" is printed in the screen.
5. You can also push image to your respository using `openEuler-image-cross-build push tag/of/your/image`. The username and password are needed during pushing.
