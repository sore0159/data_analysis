class: center, middle

## Analysis: Twitter
Eric Sorell, Dec 2016

???
Minimalism

Introduce yourself.
* Eric Sorell
* Bachelors In Science, Mathematics (abstract algebra focus)

---
class: center, middle
## Project Objective

???
Introduce the project.
* "Do Some Interesting Data Analysis"
* 30 minutes + 30min Q&A
* Expectation setting
* * Start to finish, 2 weeks
* * Focus on understanding over performing

---
layout: true
s1:.unselect
s2:.unselect
s3:.unselect
s4:.unselect
## Project Overview

.grey[
* <span class={{s1}}>Choose Data</span>
* <span class={{s2}}>Collect Data</span>
* <span class={{s3}}>Analyze Data</span>
* <span class={{s4}}>Visualize Analysis</span>
]
---
s1:select
s2:unselect
s3:unselect
s4:unselect

Twitter API Docs: 

https://dev.twitter.com/streaming/public

???
First: explain we're doing a broad overview; big picture choices

I wanted a data set that was personally motivating.  
* I have been working with English Parsing
* Lots of associated data
* Twitter is socially relevant

The particular variables considered changed over the course of the project, the main idea was that there would be _something_ interesting.
---
s2:select
s1:unselect
s3:unselect
s4:unselect

Open source Go-lang Twitter API library: 

https://github.com/dghubble/go-twitter/twitter
???
I used Go because given the short time frame of this project I always sought the tools I was most familiar with.

I quickly had a program monitoring and saving data to disk.  Not being sure what I'd use, I saved some vital stats and the word content of the tweets, to be processed later.

---
s3:select
s1:unselect
s2:unselect
s4:unselect

Open source matrix library (Golang backed by FORTRAN): 

https://github.com/gonum/matrix

???
From the start I knew I wanted to perform a Multiple Linear Regression, the famous statistical technique that drives all the fans wild.

Ideas like exploring confounded variables, controlling for effects, all were things I've encountered non-mathematically before, and so wanted to know more.

---
s4:select
s1:unselect
s2:unselect
s3:unselect

.center[![Visualization](img/visualize.png)]
???









---
layout:false
## Data Analysis of Twitter Behaviors

.biggen[
* Data Collection
* Descriptive Analysis
* Predictive Analysis
]

---
layout:true
name: data_collect
ft: 

## Data Collection
The Twitter "Sample" api provides a stream of randomly selected live tweets.<sup>1</sup> {{ content }}

.footer[
1: https://dev.twitter.com/streaming/public
{{ ft }}
]
---
The Twitter api was accessed using an open-source api-library.<sup>2</sup>
1. Filter
2. Timespan
3. Variables Measured

---

### Filter
* Tagged "English"
* Not a "retweet"
* Has geographic coordinate data
* Located within the continental United States

---

### Timespan

* Dec 2nd, 2016 to Dec XX, 2016
* Roughly 2 tweets a second
* Sporadic outages in collection
* Total collected: XXXX (>1mil)

---

### Variables Measured

Stored in CSV format on disk for later parsing/analysis
* User Numeric ID
* User Creation Date
* User Follower Count
* User Friend Count
* User Tweet Count


* Time of Tweet
* Location of Tweet
* Number of links in Tweet Content
* Copy of non-link words in Tweet Content

---
layout:false
class: middle
.center[
## Data Collection Q & A
]
```go
	filterParams := &twitter.StreamFilterParams{
		Language: []string{"en"},
		// These coords should bound the continential United States
		Locations:     []string{"-124.85,24.39,-66.88,49.38"},
		StallWarnings: twitter.Bool(true),
	}
    
```

---

## Descriptive Analysis

1. Augmenting Data
  * Location in relation to major US cities
  * Content Word Count
2. Examination Methods
  * Normalization
  * Covariance Matrices
  * Scatterplots
3. Examination Summary
  * Lots of Noisy Data

---
## Augmenting Data
### Location

Using a listing of geographic coordinates and population count of the top 1000 US cities<sup>3</sup>, I compared each tweet to the top 100 most populous cities and selected the closest city, taking note of it's population and the distance the tweet was from it's center.

### Content Word Count

Content "words" are whitespace-separated UTF-8 sequences that are not prefixed by http:// or https:// (which are counted as 'links' instead).

.footer[ 3: https://gist.github.com/Miserlou/c5cd8364bf9b2420bb29 ]
---

## Examination Methods
### Normalization
Variable measurements were normalized to (X-m)/std, and so all sample sets have a mean of 0 and std/variance of 1

### Covariance Matrices & Scatterplots
In search of linear relationships, I calculated covariance matrices of variable sets and chart all variables against each other with scatterplots using open source statistics and plotting libraries<sup>4</sup>.

.footer[ 4: https://github.com/gonum/stat ]
---
### Image Test
.center[![Sample Image](img/sample.png)]
---
class: middle
.center[
## Descriptive Analysis Q & A
]
```go
    
```
---

## Predictive Analysis

1. Multivariable Linear Regression
  * Ordinary Least Squares Fit
  * Algorithm Concerns
2. Something else
  * A
  * B

---
class: middle
.center[
## Predictive Analysis Q & A
## General Q & A
]
```go
    
```
