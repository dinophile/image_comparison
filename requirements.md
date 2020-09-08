## Given:
csv file with the following fields
| image1 | image2 |
|--------|:------:|
|aa.png  | ba.png | 

- each field contains absolute path to an image file

## Output:
- a csv file with the following fields:

| image1 | image2 | similar | elapsed |
|--------|:------:|:-------:|:-------:|
|aa.png|ba.png|0|0.006
|ab.png|bb.png|0.25|0.843|

- first two fields should match the first two fields of the given csv
- the 'similar' value should represent a score based on how similar the two images are (compare the visual elements not the binary elements)
  - so if you're given the same image in a different format it should have the same score
- elapsed is how long each comparison took in seconds
- should work on windows as well!


## considerations:
- how do you test?
- how do you teach a user how to use?
- how do you hand off to a new maintainer?
- how do you make sure the user gets the most up to date version?

