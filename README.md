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

Run this command to install: 

```bash
curl -sSL https://raw.githubusercontent.com/cjlangan/java-class-tree-cli/main/install.sh | sh
```
