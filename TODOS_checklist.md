## steps
- [ ] 1. script first 
  - [x] ~~Test comparison algo with two local images~~ turns out static path interpolation ain't so easy in go, will have to revisit this to learn what's happening! So on to fetching from t'internet!
  - [x] open a csv file 
  - [ ] line by line I need to:
      - [x] fetch the images at the path given in fields 1 and 2
      - [x] store the images for use ~~or pass them straight to~~      
      - [x] run the images through the comparison algorithm

    **Status September 7th**
      - so I can get my images in the right data type passed over to my algo! :SUCCESS:
      - however I went and chose 4 images without paying attention to size and resolution
      - so I now need to normalize the images to proceed!

      

      - [x] run the elapsed function each time the comparison algorithm runs (tentative)
      - [ ] output the results to another csv
        - [ ] image 1 path for field one, image 2 path for field 2, result of comparison, elapsed time
  - [ ] when done return a csv file with the correctly formated information

- [ ] 2. tests
    - [ ] a. unit tests for custom functions:
        - [ ] ConvertImage
        - [ ] ParseCSV
        - [ ] ImageCompare

- [ ] 3. convert to HTTP service
  - [ ] a. create webserver
  - [ ] b. create route ("/convert")
  - [ ] c. create controllers with script code 
      - [ ] a. need to be able to upload and accept a csv file
  - [ ] d. contract tests for server
      - GET returns correct data type

- [ ] 6. pipeline
    - [ ] a. connect to pipeline for any merge to master
        - [ ] necessary env variables
        - [ ] add example.env in code for maintainers and instructions on how to interact with infra - if I had a nickel for everytime this was not included in projects...
        - [ ] any build steps needed, tests to run

- [ ] 7. deploy
  - [ ] a. after pipeline is stable:
  - [ ] b. choose platform and create provisioning IoC as needed

- [ ] 8. create user documentation 
  - [ ] a. document from progress.log folder all usage requirements for users

- [ ] 9. create handoff information
  - [ ] a. How to run server/script locally
  - [ ] b. how to run tests
    - [ ] a. how to write new tests if there is a standard pattern that deviates from community norms
  - [ ] c. how to set up for deployment
    - [ ] a. any env variables needed? 
    - [ ] b. any changes to infra scripts etc
  - [ ] c. structure of how service operates. 
  - [ ] d. API documentation 

  BLUE SKY ADDS:
  1. Potentially harnessing a machine learning service/API/algorithm for further processing
  2. Front end for microservice (microfront end?) to make it even easier for users: they could just drag and drop and then download the results? Or display then download? :think:
  3. Auth? Auth. Some sort of simple token based auth would be useful, tied to user emails? Nothing with passwords or data to be stored, no one needs that hassle. Also would set up the deployment infra to only respond to internal IPs by default.
  