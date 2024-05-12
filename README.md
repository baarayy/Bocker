# Bocker

![Bocker](https://i.ibb.co/sCLfMv8/sodasweet-is-ded-by-daisydog3-debokz7.jpg)

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

3. Build the project using `make`:

   ```sh
   make all
   ```

## References

- [Part1: User and PID namespaces](http://lk4d4.darth.io/posts/unpriv1/)
- [Part2: UTS namespace (setup namespaces)](http://lk4d4.darth.io/posts/unpriv2/)
- [Part3: Mount namespace](http://lk4d4.darth.io/posts/unpriv3/)
- [Part4: Network namespace](http://lk4d4.darth.io/posts/unpriv4/)
