# Bocker
![Bocker](https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcTAbhW80JC-tv6oCCFgql-QNrRqwmG0zPoNOg&usqp=CAU)

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
