# bcryptit

This tiny app is used to generate bcrypt hashes

## Usage

Type bcryptit.exe -h for help
```
> bcryptit.exe -h
Usage of bcryptit.exe:
  -c int
        hashing cost of the bcrypt function (default 10)
  -p string
        Provide the password to hash (default "password-to-hash")
```

Example usage
```
>bcryptit.exe -p "supersecretpasswrd"
Your hashed password is:
$2a$10$L8.fk5hB6CmQwbpymuNogOw64zM/lQnCAXORfTc.JjY6kGzow7/e.
```
