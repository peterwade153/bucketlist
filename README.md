# bucketlist

### Installation.
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