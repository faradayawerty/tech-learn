
# Opening file
General syntax
```
	FILE *pointer_to_the_file = fopen("filename", "permissions");
```

Examples
```
	FILE *file_to_read = fopen("read.txt", "r");
	FILE *file_to_write = fopen("read.txt", "w");
```


# Reading from a file
```
	// like standard input and output
	fscanf(file_to_read, "format string", &var1, &var2);

	while((c = fgetc(file_to_read)) != EOF) {
		// ... do stuff with character `c`
	}
```


# Writing to a file
```
	// like standard input and output
	fprintf(file_to_write, "string");

	// write character `c` to the `file_to_write`
	fputc(c, file_to_write);
```

