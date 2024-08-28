# Bocker

![Dummy logo](https://i.ibb.co/Z28J5fH/68747470733a2f2f692e6962622e636f2f73434c664d76382f736f646173776565742d69732d6465642d62792d6461697379.jpg)

Bocker is a simple container written in Go.

It makes use of go to create isolated namespaces and chroot the default container directory to an isolated filesystem by copying /proc directory

## Installation

To run the project, follow these steps:

1. Clone the repository:

   ```sh
   git clone git@github.com:baarayy/Bocker.git
   ```

2. Navigate to the project directory:

   ```sh
   cd Bocker
   ```

3. ## Network
   set DNS resolver in container:

```sh
echo "nameserver 8.8.8.8" >> /etc/resolv.conf
```

set IP forward in host:

```sh
sysctl -w net.ipv4.ip_forward=1
```

4. Build the project using `make`:

   ```sh
   make all
   ```

## References

- [Part1: User and PID namespaces](http://lk4d4.darth.io/posts/unpriv1/)
- [Part2: UTS namespace (setup namespaces)](http://lk4d4.darth.io/posts/unpriv2/)
- [Part3: Mount namespace](http://lk4d4.darth.io/posts/unpriv3/)
- [Part4: Network namespace](http://lk4d4.darth.io/posts/unpriv4/)
