## An image comparison app in Go
##### *almost

Here I have the beginnings of an application to compare two `.png` images for similarity.

I say beginnings of because as you can guess: it's not quite working. It's not nearly as polished and finished as I had hoped. Choosing a language that's a 180 from what I'm used to working in made me lose time. The easiest part of using golang for me is setting up a webserver, one of my later tasks. The script to compare images took the bulk of my time. 

I can get the script most of the way there with `go run main.go`, with the last hurdle being resizing the images to match in size and possibly resolutions? Need to investigate that! I'm getting stuck on `image bounds not equal` and that's pretty cool! I made it a lot farther than I had thought! :D

Why did I choose Golang? I do love a challenge. But also if given a large number of images to compare, go would be faster to cruch through them vs something like javascript. Also since getting this script working on multiple operating systems could be a possibility I can compile and pass on a program to different OS users.

My ultimate goal though was to create this as a simple API to be used by anyone (within an organization!). 

I had planned one simple endpoint, `/convert` to accept a csv file in the format shown in `requirements.md`. A stretch goal would be to make a simple front end for folks not familiar with `curl`.

I've kept track of high level milestones in `TODOS_checklist.md` and a working progress journal in `progress_log`.

### So, not successful...yet
But I've made a lot of progress! I'm pretty proud I made it almost to the end of a working script in a language I'm still very new with. And the opportunity to try to build out something new helped me learn more about images in go than I ever thought possible!

## My one regret though: 
Is that I haven't yet added tests. Once I get my script working and returning a satisfactory csv I really do need to go back and write unit tests 