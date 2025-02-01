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

    // Read the contents of the directory
    files, err := os.ReadDir(directory);
    if err != nil {
        fmt.Println("Error reading directory.");
        return;
    }

    // Hashmap to store all base nodes
    graph := map[string]*Node{};

    // Iterate over the files in the directory
    for _, file := range files {
        if !file.IsDir() {
            //fmt.Printf("Found file %s\n", file.Name());
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

                // proceed if there is a class
                if class != "" {
                    //fmt.Printf("Class: %s\tParent: %s\n", class, parent);
                    var child *Node;

                    // Create and add the child if doesn't exist
                    if graph[class] == nil {
                        child = &Node{class: class, hasParent: false, subs: []*Node{}};
                        graph[class] = child
                    } else {
                        child = graph[class];
                    }
                    
                    if parent != "" {
                        // Make child have parent
                        child.hasParent = true;

                        // Add child to parent
                        if graph[parent] == nil {
                            super := &Node{class: parent, hasParent: false, subs: []*Node{}};
                            graph[parent] = super;
                        }
                        graph[parent].subs = append(graph[parent].subs, child);
                    }
                }
            }
        }
    }
    // Print the nodes
    for _, node := range graph {
        if(!node.hasParent) {
            node.printNode(0);
        }
    }
}

// Print node, then its children
func (n *Node) printNode(indent int) {
    fmt.Print("─");
    fmt.Println(n.class);

    // Print children
    i := 0;
    for _, node := range n.subs {
        // Spacing
        for j := 0; j < indent + 1; j++ {
            fmt.Print(" ");
        }

        // determine if last
        if i == len(n.subs) - 1 {
            fmt.Print("└");
        } else {
            fmt.Print("├");
        }
        node.printNode(indent + 1);
        i++;
    }
}

// A node has a class, maybe a parent, and a list of subclasses
type Node struct {
    class string;
    hasParent bool;
    subs []*Node;
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


