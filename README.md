[![Build Status](https://travis-ci.org/peterwade153/bucketlist.svg?branch=master)](https://travis-ci.org/peterwade153/bucketlist)

# bucketlist go-lang trial

### Installation.
Should have Go Installed
Clone the repository.
<pre>
git clone https://github.com/peterwade153/bucketlist.git
</pre>

### Running application
This uses memory and quick to set up and use.
Change directory into bucketlist then
<pre>
$ go run main.go
</pre>

API endpoint can be accessed. Via http://localhost:5000/

### Endpoints

Request |       Endpoints                 |       Functionality
--------|---------------------------------|--------------------------------
POST    |  /items                         |        Add item to bucketlist requires( id, title, description)
PUT     |  /items/id                      |        Edit an item in bucketlist requires( title, description)
DELETE  |  /items/id                      |        Delete an item in a bucketlist
GET     |  /items/id                      |        Returns an item in a bucketlist
GET     |  /items                         |        Retuens all items in a bucketlist