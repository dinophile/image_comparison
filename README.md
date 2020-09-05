## steps
- [ ] 1. script first 
  - [ ] Test comparison algo with two local images
  **STATUS SEPT 4 2020**
    - hiccuping on opening a local image file, having issues with getting the path to the image correct, but progress made in `main.go`
    - ignore `scratchpad.go` just using that as a dumping ground for now!
  - [ ] open a csv file 
  - [ ] line by line I need to:
      - [ ] fetch the images at the path given in fields 1 and 2
      - [ ] store the images for use? or pass them straight to:
      - [ ] run the images through the comparison algorithm
      - [ ] run the elapsed function each time the comparison algorithm runs
      - [ ] output the results to another csv
        - [ ] image 1 path for field one, image 2 path for field 2, result of comparison, elapsed time
  - [ ] when done return a csv file with the correctly formated information

- [ ] 2. tests
    - [ ] a. tests for custom functions:
        - [ ] ConvertImage
        - [ ] FastCompare

- [ ] 3. convert to HTTP service
  - [ ] a. create webserver
  - [ ] b. create route ("/convert")
  - [ ] c. create controllers with script code 
      - [ ] a. need to be able to upload and accept a csv file
  
- [ ] 4. tests for webserver
    - [ ] a. create contract tests for endpoint
  - [ ] b. create any unit tests needed for extra custom functions

- [ ] 5. pipeline
    - [ ] a. conenct to pipeline for any merge to master
        - [ ] any build steps needed, tests to run

- [ ] 6. deploy
  - [ ] a. after pipeline is stable:
  - [ ] b. choose platform and create provisioning IoC as needed

- [ ] 7. create user documentation 
  - [ ] a. document from progress.log folder all usage requirements for users

- [ ] 8. create handoff information
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
  