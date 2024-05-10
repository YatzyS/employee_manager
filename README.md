# Employee Manager

This service manages the employees database. Currently it is creating an in memory data store for the employees, where it stores the Id, Name, Position and Salary of the employees. It provides CRUD APIs for the DB. It also provides a List API with Pagination.

### APIs

---

***Create Employee***

----
Returns the newly created employee.

* **URL**

  /v1/create

* **Method:**

  `POST`


* **Data Params**

  `name`, `email`,`age`

* **Success Response:**

    * **Code:** 200 <br />
      **Content:** `{
      "id": 1,
      "name": "test",
      "position": "developer",
      "salary": 100000
      }`

* **Error Response:**

    * **Code:** 400 BAD REQUEST <br />

  OR

    * **Code:** 500 INTERNAL SERVER ERROR <br />

* **Sample Curl:**
    ```
    curl --location --request POST 'localhost:8080/v1/create' \
    --header 'Content-Type: application/json' \
    --data-raw '{
    "name": "test10",
    "position": "developer",
    "salary": 100000
    }'
    ```
----

***Update Employee***

----
Used to update the employee data.

* **URL**

  /v1/update

* **Method:**

  `POST`


* **Data Params**

  `id` (Mandatory)
  `name` OR `position` OR `salary` (Any, all or multiple can be used as per the requirmenet)

* **Success Response:**

    * **Code:** 200 <br />

* **Error Response:**

    * **Code:** 400 BAD REQUEST <br />

  OR

    * **Code:** 500 INTERNAL SERVER ERROR <br />

* **Sample Curl:**
    ```
    curl --location --request POST 'localhost:8080/v1/update' \
    --header 'Content-Type: application/json' \
    --data-raw '{
    "id":1,
    "Name":"Test",
    "Position": "Senior Dev",
    "Salary":1000000
    }'
    ```
----

***Delete Employee***

----
Used to delete a specific employee.

* **URL**

  /v1/delete

* **Method:**

  `DELETE`


* **Query Params**

    * `id` (Mandatory)

* **Success Response:**

    * **Code:** 200 <br />

* **Error Response:**

    * **Code:** 400 BAD REQUEST <br />

  OR

    * **Code:** 500 INTERNAL SERVER ERROR <br />

* **Sample Curl:**
    ```
   curl --location --request DELETE 'localhost:8080/v1/delete?id=2'
    ```
----

***List Employees***

----
List all the employees.

* **URL**

  /v1/list

* **Method:**

  `GET`


* **Query Params**
    *  page
    *  pageSize

* **Success Response:**

    * **Code:** 200 <br />
    * **Content:** `[
      {
      "id": 1,
      "name": "test10",
      "position": "developer",
      "salary": 100000
      },
      {
      "id": 3,
      "name": "test10",
      "position": "developer",
      "salary": 100000
      },
      {
      "id": 4,
      "name": "test10",
      "position": "developer",
      "salary": 100000
      }
      ]`

* **Error Response:**

    * **Code:** 500 INTERNAL SERVER ERROR <br />

* **Sample Curl:**
    ```
    curl --location --request GET 'localhost:8080/v1/list?page=1&pageSize=3'
    ```
----

***Get Employee By Id***

----
Used to get data of specific employee.

* **URL**

  /v1/get

* **Method:**

  `GET`


* **Query Params**

  `id` (Mandatory)

* **Success Response:**

    * **Code:** 200 <br />
    * **Content:** `{
      "id": 1,
      "name": "test",
      "position": "developer",
      "salary": 100000
      }`

* **Error Response:**

    * **Code:** 400 BAD REQUEST <br />

  OR

    * **Code:** 500 INTERNAL SERVER ERROR <br />

* **Sample Curl:**
    ```
    curl --location --request GET 'localhost:8080/v1/get?id=1'
    ```
----

### Test Coverage
The repo layer is covered with unit test. It can be seen by running the following command in the /internal/repo/employee folder
```
go test -cover
```
The current coverage is ~98% for that folder