# Writing a Go module that calls an external C++ library

Let's assume you've got go installed. 

### Step 1: Create a go app which will be your module

   ```bash
   mkdir mymodule
   cd mymodule
   go mod init mymodule
   ```

### Step 2: Create the C++ Library and header file

1. **Write the C++ Application Code**: Create a C++ source file, say `mylib.cpp`, with the functions you want to call from Go.

   ```cpp
   // mylib.cpp
   #include <iostream>

   extern "C" {
       void sayHello() {
           std::cout << "Hello from C++!" << std::endl;
       }

       int add(int a, int b) {
           return a + b;
       }
   }
   ```

2. **Create a C Header File**: Create a C header file, say `mylib.h`, to declare the functions.

   ```c
   // mylib.h
   #ifndef MYLIB_H
   #define MYLIB_H

   void sayHello();
   int add(int a, int b);

   #endif
   ```


3. **Compile the C++ Library**: Compile the C++ code into a shared library.

   ```bash
   g++ -shared -o libmylib.so -fPIC mylib.cpp
   ```
   - The `-shared` flag tells the compiler to create a shared library.
   - The `-fPIC` flag generates position-independent code, necessary for shared libraries.
   - The output is `libmylib.so`, the shared library.


### Step 3: Write the Go Code

1. **Create a Go File**: Create a Go source file, say `main.go`, and use cgo to call the C++ library.

   ```go
   // main.go
   package main

   /*
   #cgo LDFLAGS: -L. -lmylib
   #include "mylib.h"
   */
   import "C"

   import "fmt"

   func main() {
       C.sayHello()

       result := C.add(1, 2)
       fmt.Printf("Result of add: %d\n", int(result))
   }
   ```

   - The `#cgo LDFLAGS: -L. -lmylib` directive tells cgo to link against the `libmylib.so` library in the current directory.
   - The `#include "mylib.h"` directive includes the C header file with function declarations.
   - The `C` package is used to call C functions.
   - The `main` function calls the `sayHello` and `add` functions from the C++ library.



### Step 4: Build and Run the Go Program

1. **Set Up the Environment**: Ensure that the shared library and header file are in the same directory as the Go code.


2. **Build the Go Program**: Use the `go build` command to build your Go program. Make sure the shared library is in the library path.

   ```bash
   go build
   ```

3. **Run the Go Program**: Execute the built program.

   ```bash
   ./mymodule
   ```

If everything is set up correctly, you should see the output:

```
Hello from C++!
Result of add: 3
```


