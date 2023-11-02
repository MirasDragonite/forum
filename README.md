## Forum
* `mkabyken`
* `etolmach`
* `Adilkhan2001_alem` 

### Description:
This project involves creating a web forum with features like, post categorization, liking and disliking posts and comments, and post filtering. The data is managed using SQLite, a popular embedded database, and the project encourages optimizing performance with an entity relationship diagram. User authentication allows registration, login sessions with cookie management, and optional UUID usage. Users can create posts and comments, associate categories with posts, and engage with likes and dislikes. A filter system allows users to sort posts by categories.


### Usage
Clone the repository:
```
git clone git@git.01.alem.school:mkabyken/forum.git
```

#### Run with docker

```
docker build -t forum .
```
```
docker run -dp 8000:8000 forum
```
#### Run without docker

Run a program:
```
go run ./cmd/main.go
```
Open the link in browser
```
localhost:8000
```
 