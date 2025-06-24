# Workshop

The objective of the workshop is to try some of the different tools
Cursor has to facilitate onboarding and collaborating in a codebase.

We will use this exercism cli and the exercises in exercism.org to
use some of these tools.

## Setup

> [!IMPORTANT]
> Use Cursor to help with the setup.

During this workshop we will be using Exercism to test the workflows, to do so you need to:

1. Create an account in exercism.org
2. Get your account token (keep it handy)
   Profile -> Settings -> API / CLI -> API/CLI Token
3. Follow [these instructions](https://github.com/mini024/cli/blob/ewit/workshop/CONTRIBUTING.md#manual-testing-against-exercism) to set up manual testing
4. Run ./testercism configure --token=YOUR_TOKEN

> [!IMPORTANT]
> Use _./testercism_ to run the commands and test your changes.

## Activities

## Activity 1: Understand the Code

Explore the different commands that the exercism cli has to facilitate
submiting exercises to the website.

1. What languages does the code base use? Does it use a specific testing library?
2. What does the download command do? What attributes does it need? 
3. How do you run the tests for a command? Are any tests currently failing? 

## Activity 2: Remove a Command

We have decide to remove the prepare command because maintances is high and usage
is low. Create a branch and remove the required code so this command is
not available anymore.

If you try to run the command what happens? 
How do we confirm that the command was removed correctly? 

> Note: there is no need to push these changes.

## Activity 3: Implement a feature

Define a clear set of requirements and using Cursor implement the following feature.

### Feature
We want to implement a command "open" that will open the url to an exercise.

### Design
This command will take the exercise directory as an argument, find the
url in the metadata file and open it.

### Downloading an Exercise
To try downloading an exercise and seeing the structure of the directory:
1. Got to https://exercism.org/tracks
2. Select and join a track you would like to explore
3. Choose an exercise
4. Find the "Work Locally (CLI)" section on the left and get the track and exercise.
5. Download the exercise using `./testercism download --track= --exercise=`

After you download the project you can explore what it contains and use this
to get the information needed to open the url. 

## Activity 3.2: Compare

The Open command is an [existing command](https://github.com/exercism/cli/blob/main/cmd/open.go) in the main branch. 

Compare your solution with the implemented one, does the code look different? 

## Activity 3.3: Implement the tests for open command

We want to make sure the open command works as expected, implement some tests for the command.

## Challenge

Go to Exercism.org and find a path that you are interested in and solve some challenges using Cursor!
