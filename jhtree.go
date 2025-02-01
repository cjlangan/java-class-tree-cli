package main 

import (
    "fmt"
    "os"
    "strings"
) 

func main() {
    // Take in directory as input
    if len(os.Args) < 2 {
        fmt.Println("Expected a directory to analyze.");
        return;
    }
    directory := os.Args[1];
    fmt.Println("Analyzing directory..");

    // Read the contents of the directory
    files, err := os.ReadDir(directory);
    if err != nil {
        fmt.Println("Error reading directory.");
        return;
    }

    // Iterate over the files in the directory
    for _, file := range files {
        if !file.IsDir() {
            fmt.Printf("Found file %s\n", file.Name());
            nameLen := len(file.Name());

            // Last 5 chars of file name
            lastPart := file.Name()[nameLen - 5:];

            // Check that the file is a .java file
            if lastPart == ".java" {

                // get the contents of the file
                content, err := os.ReadFile(file.Name());

                if err != nil {
                    fmt.Println("Error reading file");
                    continue;
                }

                // Determine the class and parent
                lines := strings.Split(string(content), "\n");
                class, parent := getClassAndParent(lines);
                fmt.Printf("Class: %s\tParent: %s\n", class, parent);

                // Add them to hierarchy
            }
        }
    }
}

// determines the class and parent class of a java file
func getClassAndParent(lines []string) (string, string) {
    i := 0;
    done := false;
    class := "";
    parent := "";

    // Find class line
    for !done {
        if strings.Contains(lines[i], "class") {
            done = true;
        } else {
            i++;
        }
    }

    if done {
        words := strings.Split(lines[i], " ");

        // find names of classes
        for i := 0; i < len(words); i++ {
            if words[i] == "class" {
                class = words[i + 1];
            }
            if words[i] == "extends" {
                parent = strings.Trim(words[i + 1], "{}");
            }
        }
    } else {
        fmt.Println("keyword class not found.");
    }

    return class, parent;
}


