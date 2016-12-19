# Object Version Reconstruction

Object Version Reconstruction is a web application that allows you to upload a CSV file containing the changes in properties of different objects at different times, and then allows you to query the state of an object at a particular timestamp.

This is what the table would look like:

object_id | object_type | timestamp | object_changes
:-------: | :---------: | :--------: | :------------
 1        |  ObjectA    |  412351252 | {"property1": "value", "property3": "value"}
 1        |  ObjectB    |  456662343 | {"property1": "value"}
 1        |  ObjectA    |  467765765 | {"property1": "altered value", "property2": "value"}
 2        |  ObjectA    |  451232123 | {"property2": "value"}
...       |  ...        |  ...       | ...

The CSV columns are:

 - **object_id:** is a unique identifier per-object type.
 - **object_type:** denotes the object type.
 - **timestamp:** needs no explaination
 - **object_changes:** the properties changed for specified object at **timestamp**.

## How to use the application:

You can visit [demo](http://13.76.208.65:9090/) here.
The home page allows you to upload a CSV file. Take a look at `test_resources/test.csv` for the ideal format of the CSV file (note that the timestamp is Unix timestamp format, and the object_changes is in JSON format).

Each time a new csv file is uploaded, the database is flushed and data from the csv file is inserted into the database. Once the file is uploaded you are directed to a page which displays all the object changes, and allows you to filter by `object type`, `object type and object id` and `object type, object id and timestamp`.


## Setting up the dev environment

1. [Install Go](https://golang.org/doc/install)

2. [Install MySQL](http://dev.mysql.com/doc/refman/5.7/en/installing.html)

3. Clone this project into your workspace by running:

	```
	git clone https://github.com/anand-sundaram/object-version-reconstruction.git
	```

4. Run the following commands to import the external packages:

	```
	go get github.com/go-sql-driver/mysql
	go get github.com/gorilla/mux
	```

5. Update the database details in dbconfig.sample.go and rename the file to dbconfig.go

6. Run the following command from the project root folder:

	```
	go build
	```
7. The tests can be run with the following command from the project root folder:

	```
	go test ./...
	```


8. Start the application by running the following from the project root folder

	```
	./object-version-reconstruction
	```

The application will now be running at [http://localhost:9090](http://localhost:9090).

## Stack

This application is built using Go. As this was my first time using Go I started out by looking for a web framework to use, but based on the recommendations I found online, I chose not to use any existing framework and instead make use of the language's rich libraries. The data from the CSV file is stored in a MySQL database.
