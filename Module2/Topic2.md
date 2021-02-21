## Packages


1. A library , module or namespace in any other language is called a package. Packages are a way to structure code. A program is contsructed as a package which may use functionalities from other packages. A packages is often abbreviated as "pkg"

2. Every Go file belongs to one package whereas one package can comprise many different Go files. Hence, the filename and the package name are usually not the same. The package to which the code belongs must be indicated on the first line.

3. The `import` keyword is used to import a package within a different package. The package names are enclosed within double quotes (" ")

4. Visibility rule :

    1. Packages export their code objects to code outside of the package according to the following rules enforced by the compiler.

    2. When the identifier starts with an uppercase letter, like, Group1, then the 'object' with this identifier is visible in code outside the package and it is said to be exported. Identifiers that start with a lowercase alphabet are not visible outside the package.

5. Go has a motto known as no unncessesary code, So, importing a package which is not used results in a build error.