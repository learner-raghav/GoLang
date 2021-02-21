1. Strings are a sequence of UTF-8 characters (the 1-byte ASCII code is used when possible, and a 2-4 buyte UTF-8 code whenever necessary)

2. UTF-8 is the most widely used encoding. It is the standard for encoding for text files, XML Files and JSON Strings. 

3. With the string datatype, we can reserve 4 bytes for characters, but Go is intelligent enough that it will only reserve 1 byte if the string is only an ASCII Character.

4. Strings in Go are immutable, IN other words strings are immutable arrays of bytes.

5. The length of a string in Go can be calculates as len(string) and we can concatenate two strings using s as 
    s := s1 + s2, this is same as presnet in various other languages like Python, Java etc.

