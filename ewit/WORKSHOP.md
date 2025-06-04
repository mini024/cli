# Workshop

The objective of the workshop is to try some of the different tools
Cursor has to facilitate onboarding and collaborating in a codebase.

We will use this exercism cli and the exercises in exercism.org to
use some of these tools.

## Setup

During this workshop we will be using Exercism, to do so you need to:

1. Create an account in exercism.org
2. Get your account token (keep it handy)
3. Sign up for a language track online
4. Follow [these instructions](../CONTRIBUTING.md#manual-testing-against-exercism) to set up manual testing
5. Run `./testercism configure --token=YOUR_TOKEN`

**Note**
Cursor might suggest installing the command-line client. How can we ensure it uses this local version of the CLI?

## Activities

## Part 1: Understand the Code

Explore the different commands that the exercism cli has to facilitate
submiting exercises to the website.

1. What is the order in which the commands are used?
2. What happens when you submit an exercise?
3. How does it link a submission to a solution?

## Part 2: Remove a Command

We have decide to remove the prepare command because it is not
really useful. Create a branch and remove the required code so this command is
not available anymore.

How do we confirm that the command was removed correctly?

## Part 3: Implement a feature

Define a clear set of requirements and using Cursor implement the following feature.

We want to implement a command "open" that will open the url to an exercise.

When you select an exercise to work on you can use the download command to
pull the repo where the solution can be implemented. This repo contains a
.exercism/metadate.json file that has the URL where the exercise can be
found on the website.

This command will take the exercise directory as an argument, find the
url in the metadata file and open it.

### Downloading an Exercise

1. Go to exercism.org
2. Select a Track and chose an exercise
3. Find the "Work locally option"
4. Update the command to use ./testercise instead of exercism
5. Download the exercise

## Part 4: Compare

The Open command is an existing command in main. Compare your solution
with the implemented one, what is different?

We want to make sure the open command works as expected, write a test open_test.go.
