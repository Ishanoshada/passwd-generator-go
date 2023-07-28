# Password Generator

![Go](https://img.shields.io/badge/Go-Programming%20Language-blue)


**Password Generator** is a simple command-line tool written in Go that generates random passwords containing letters (both uppercase and lowercase), numbers, and special characters. It's a handy utility to create strong and secure passwords for various purposes, such as online accounts, applications, or anything that requires a secure passphrase.

## Folder Structure

The folder structure for the Password Generator project is organized as follows:

```
passwd_generator_go/
|-- main.go
|-- go.mod
|-- README.md
```


## How to Use

1. Clone the repository to your local machine:

   ```bash
   git clone https://github.com/ishanoshada/passwdgenerator_go.git
   ```

2. Navigate to the project directory:

   ```bash
   cd passwd_generator_go
   ```

3. Build and run the `main.go` file using Go:

   ```bash
   go run main.go
   ```

4. You will be prompted to enter the desired length for the password.

5. The generated password will be displayed on the console.

## Examples

Here are some examples of using the Password Generator:

1. Generating a 12-character password:

   ```bash
   Enter password length: 12
   Generated Password: Xq1@zv9Ld^W2
   ```

2. Generating a 20-character password:

   ```bash
   Enter password length: 20
   Generated Password: p7N$RfG@U0Xo5Q4ma3i9
   ```

3. Generating a 8-character password:

   ```bash
   Enter password length: 8
   Generated Password: cF7L6@Kx
   ```

## Customize

You can easily customize the character set used for generating passwords by modifying the `charset` constant in the `main.go` file. Additionally, you can add or remove special characters or adjust the distribution of characters to suit your specific requirements.

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

## Author

This Password Generator tool was created by [ishan oshada](https://github.com/ishanoshada). Feel free to contribute to the project by submitting issues or pull requests.

## Acknowledgments

Special thanks to the Go programming language community for providing an excellent programming language that makes it easy to develop efficient tools like this password generator.

## Support

If you find this tool helpful, consider giving it a ⭐️ on GitHub and sharing it with others! If you encounter any issues or have suggestions for improvements, please [open an issue](https://github.com/ishanoshada/passwd_generator_go/issues).

