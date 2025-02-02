# jhtree 

#### A tool to print you java class hierarchy 

## Usage 

Simple run `jhtree <directory>`. Such as `jhtree .` for the current directory:

```
MealTracker
Item
├─Composite
└─Single
   ├─Fruit
   ├─Substitute
   │  ├─DairyMeatGrain
   │  │  ├─Dairy
   │  │  ├─Grain
   │  │  └─Meat
   │  └─Other
   └─Vegetable
      └─Carrot
Library
├─Foods
└─Profiles
Node
└─User
```

## Installation

### Linux

Run this command to install: 

```bash
curl -sSL https://raw.githubusercontent.com/cjlangan/java-class-tree-cli/main/install.sh | sh
```

### Windows 

Download `jhtree.exe` from the releases page.

Move it to a folder in your PATH, such as `C:\Program Files\jhtree\`

### MacOS

Download `jhtree-mac` from the releases page. 

Open a terminal and run: 

```bash
chmod +x jhtree-mac 
sudo mv jhtree-mac /usr/local/bin/jhtree
```
