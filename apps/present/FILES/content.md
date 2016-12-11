class: center, middle

## Analysis: Twitter
Eric Sorell, Dec 2016

Test
???
Testing Speaker Notes

---
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
ft: <br>2: https://github.com/dghubble/go-twitter/twitter
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
