# Regen Blockchain Engineer Challenge

My implementation of adding commenting functionality to the cosmos blog demo.

## Summary

  - [Getting Started](#getting-started)
  - [Runing the tests](#running-the-tests)
  - [My Approach](#my-approach)
  - [Known Issues](#known-issues)


## Getting Started

Clone the repo to your machine

### Prerequisites

Go 1.15+

## Running the tests

CD into the x/blog/client/cli directory and run the command:

`go test`

### What these tests do

TestCreateComment

    Creates a dummy post and create a comment with the posts ID.

TestAllComments

    Creates a post, create 2 comments, then query all comments with the dummy post's ID.

## My Approach

My git history doesn't really tell the full approach so here's a better version

1. Started with the proto files to define commenting protobuf service
2. wrote the tx and query methods
3. implemented methods on msg_server and query_server
4. Wrote tests for these methods
5. 

## Known Issues
The errors from msg_server don't actually propigate into the tests. 
I tried returning only an error in both CreatePost and CreateComment and neither showed up in the cli tests. Each just resulted in the transaction response code not being 0.

The way I query comments is kinda hacky. I assume theres a cleaner way
to query the KVStore for a specific Key/Value pair. 