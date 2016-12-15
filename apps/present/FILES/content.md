class: center, middle

## Analysis: Twitter
Eric Sorell, Dec 2016

???

Hello, my name is Eric Sorell, and I'm here to present on my Data Analysis project.   This will be about 30 minutes, with Q&A afterward, but please feel free to ask questions at any time.

---
class: center, middle
## Project Objective

???

I have a bachelors in mathematics with a focus on abstract algebra, so while I have a foundation of probability theory, statistical modeling is all new material.  This was a two week project from first concept to final presentation.

I hope with this project to demonstrate my ability to learn whatever material is required for whatever tasks are needed.  As a major job role is to be the explanation of our work to clients, I have focused on using techniques and tools I fully understand and can explain, avoiding more complex tools that I _couldn't_ explain.

I do not intend to present to _you_ as though you were clients, but if you would like any aspect of the project explained as though you were, please let me know.

---
layout: true

## Initial Approach

.topbox[
* Twitter is interesting, I'll get data from there
* Multiple Linear Regressions are interesting, I'll do one of those
* Scatterplots are interesting, I'll use them
]

---
class: center, middle

???

The structure of this project was determined by my initial approach to it.  "Do some interesting Data Analysis", was my instruction.  I selected my data collection, analysis, and presentation on those terms.

---

Twitter provides lots of data to developers with an API accessible via various open source libraries.

https://dev.twitter.com/streaming/public
<br>
https://github.com/dghubble/go-twitter/twitter

Tweet Filter used:
* Not a Retweet
* English language
* Geographic coordinates provided
* Located within the continental US

???

Knowing I wanted to use data on the properties of tweets, I quickly set up a program to monitor and record a constant stream of tweets.  Twitter provides a few developer APIs for this purpose, broadcasting a "random sample" of tweets as they happen.

I had this program running at all hours of the day, not wishing to bias any particular time period.  There were one or two outages of a few hours, but over a weeks worth of content was recorded; over 1.5 million tweets.

I used an initial filter of my own, knowing that I wanted to use geographic location and content as potential data to analyze.  

---

Information monitored and recorded:
* Tweet geographic coordinates
* Tweet post date
* Number of links in tweet content
* List of words in tweet content
* Poster unique numeric ID
* Poster creation date
* Poster follower count
* Poster Tweet count

???

I recorded a wide range of properties of each tweet, but by no means everything.  Twitter provides a _lot_ of data attached to each tweet.  I tried to pick what I thought would be a good base to allow for more complicated processing and analysis down the line.

---
Initial Goals for Multiple Linear Regression:
* Find some relationships between properties of the data
* Demonstrate 'controlling for confounding variables'
* Make some predictions

???

There are lots of options for computing multiple linear regressions, from R to Excel spreadsheets.  I used a golang matrix library to solve the least-squares using QR decomposition, with the heavy computation backed by calls to FORTRAN for speed.  

In the case of linear regression, I expect "just have R do it" is a fine approach, but I wanted to keep the calculation as part of a larger program of data manipulation and transfer.  I saw value in understanding the constraints on the process different algorithms impose.

---
Examination of different properties of the data for linear relationships focused on:
* Number of followers
<br> <br>
* Tweet count
* Number of words used
* Number of links used
* Distance from nearest major US city (top 100)
* Population of nearest major US city

???

I quickly assembled a set of measurable and possibly related numbers so I could start building the analysis apparatus around _something_.  

For the dependent variable, I decided on the number of followers a user had at the time of posting.  As a quality of the post, this can be thought of as trying to predict how many eyes are going to see a post by relating other qualities of the post.

The dependent variables were the amount of previous posts by the author, the number of words used in the post, the number of links used, and two qualities based on the geographical coordinates of the post.

From a list of the top 100 major US cities I calculated the closest city to the post, then recorded the distance to and the population of that city.  

---
Initial scatterplots highlighting problems

.fitimg[
![Scatter 1](scatter1.png)
![Scatter 2](scatter2.png)
![Scatter 3](scatter3.png)
]

???

Once I had some experimental data, I started exploring plots.  Scatterplots proved helpful in identifying issues with the data early on.  Here are three example plots.

First, the data had huge outliers in both popularity and post count.  Most of the data here is clumped up near the axis, with barely visible data points taking up all the rest of the space.

Second, nonsensical data.  Plots of population showed all data taking only two values, revealing that all posts were closest only to Miami or Honolulu.  This quickly led to a discovery of a bug in distance calculations: Twitter uses Longitude, Latitude for its coordinates instead of the normal Latitude, Longitude.

Third, plots showed some issues with clumping of data on variables such as word count and link count.  

---
layout: true

## Development of Analysis

.topbox[
* Processing the data
* Examining model operation
* Improving visual feedback
]

---
class: center, middle

???

Once some data was collected, some analysis were run and some plots made, I had something that worked!  The next step was to make it all _useful_.

I had initially hoped to do some more interesting content analysis (wordnet), but quickly decided that would be a lot of work without adding much analytical depth.

I after some trial models with different variable sets, I added "Age of posting account at time of post" to the independent variables.

---
title: hist_before
![Hist Followers Before](hist_bad_followers.png)
![Hist Tweets Before](hist_bad_tweet.png)

???

To better explore the variables as I tried including different properties in the model, I used histograms of the variable distributions.  Shown here are histograms that show the outlier problem in more detail.

For both post count and popularity, basically all of the data is very close to the mean.  Such outliers were obscuring any relationships in the bulk of the data in my regression fits, and so in order to fit these variables I sought some ways to process the data.

Applying domain-specific knowledge, my estimate was that accounts with extraordinary post counts were likely different from normal accounts by being non-human.  Automated weather reports, job postings, and sales messages behave quite differently from someone who just wants to share a picture of his breakfast.

Extraordinary follower counts, while likely the result of human activities, were likely due to non-twitter related factors.  Wanting to use the data I had to analyze the relationships in twitter behaviors of normal humans, I attempted to filter my data.

Apologies to Stephen Fry and his 13 million followers, but >240 is too many standard deviations.

---
title: hist_after
![Hist Followers After](hist_good_followers.png)
![Hist Tweets After](hist_good_tweet.png)

???

Testing showed meaningful distribution change required filtering out accounts with over 20,000 tweets or 2,000 followers.  The filtered distributions are shown here.  More sophisticated filtering might have used tweets/day average instead and numerical comparisons of the distributions.

The sample size goes from 1.1mil samples to 850k samples, which seems fine.

---
![Hist Age Before](hist_bad_age.png)
![Hist Age After](hist_good_age.png)

???

As a check, I looked at histograms of the other properties to see if their distributions were drastically changed by this filter and saw no huge changes to the general shapes.  Shown here is "Age of account"

---
![Scatter Before](scatter_before.png)
![Scatter After](scatter_after.png)

???

Before and after scatterplots of the post count against the follower count.  No linear relations jump out, but the data is no longer obviously dominated by a minority of the sample.

This chart is overplotted, a problem I tackle later.

---

```r
Initial R^2 = 0.002

"Humans Filter" R^2 = 0.174

Power Law R^2 = 0.296
```

???

In addition to using histograms to analyze the effect of filters on sample distributions, I also examined the effect these changes to the sample sets had on the regressions.

I used the R Squared of the regressions to get a general idea of the fit while keeping an eye on the graphs of the lines over the scatterplots to watch for problematic behavior.

The R Squared of the initial variable set was very low, but with a "humans filter" it shot up to .174  This is still a low percentage of variability explained by the model, but given the noisy nature of twitter behavior, I'm happy with this improvement.

Testing filtering a sample to only include each account once led to a drastic reduction of sample size by three fourths.  In the end, I wanted to keep this model a prediction of tweets, not users, and so did not use such a filter.

Further exploration led to even better results in both fits and data distributions.

(MeanSqResiduals 0.998   0.826   0.704)

---

![Hist Followers Power](hist_power_followers.png)
![Hist Tweets Power](hist_power_tweet.png)

???

Followers and Post Count specifically have much more normal distributions when we take their logarithm.  

###EXPLANATION OF LOG-NORMAL DISTRIBUTION GOES HERE

I opted to use this transformation instead of filtering the data because I did not have a reasonable method to make a non-arbitrary filter for the data.  Perhaps robots really should be removed from this analysis (I suspect that bump at the end of the post count chart is due to robot activity), but I can't yet do so without losing (in a biased fashion) a lot of real data.

---

* Number of followers  --> ln(x)
* Tweet count          --> ln(x)
* Number of words used
* Number of links used
* Distance from nearest major US city (top 100)    --> Cut!
* Population of nearest major US city --> ln(x)
* Age of Account --> Added

???

So, the real development of the model is here, in what variables we are including in the regression.  The distance variable turned out to have a low correlation with popularity, and it's inclusion/removal did not change the other coefficients much, so for the sake of simplicity it was removed.

The bulk of my work here has been building the tooling, for data acquisition, processing, modeling, and presentation.  I think with all that tooling now in place, were I to put an equal amount of work in, most of it would be in searching out more varied sets of variables among the data.

---
layout: false
class: center, middle
![Followers Vs Residuals](fnl_resids.png)

???

To solve overplotting in my scatterplots, I set a very small dot size, and set each data point to be very transparent (alpha of 3!).  The exact alpha to set seems to be very dependent on how clustered the data is: for 1.6mil points it had to be very low.

Shown here is an example of my residual checks for anomalies.  Looks pretty linear with a good amount of variation.

For data plots, I include regression lines and wanted to include confidence intervals but had some technical issues with that. 

---
layout: true

## Summary of the Results
---
class: middle, center

---

```go
LnFollowers Mean: 6.011637, STD: 1.398575
Links Mean: 0.997370, STD: 0.389061
Words Mean: 13.580750, STD: 3.966365
LnTweetCount Mean: 7.941631, STD: 2.408930
LnPopulation Mean: 13.502904, STD: 1.033709
Age(days) Mean: 2101.801189, STD: 828.125819

LnFollowers = (0.000000) + (0.160706)[Links] + (0.110461)[Words] +
(0.439769)[LnTweetCount] + (0.040290)[LnPopulation] + 
(0.327982)[Age(days)]

MeanSqError: 0.704020, MeanSquareResiduals: 0.704017, R^2: 0.295982
```

???
So here's the linear regression equation, with some associated stats.

Location statistics were not good predictors of popularity.  Tweetcount was the best, age of account was close second, and both using links and wordiness were close to each other as mild predictors.

---

```r
// X1 = ("LnFollowers")
    Estimate   Std. Error   t value Pr(>|t|)
X2 0.1607064 0.0006834654 235.13464        0 // ("Links")
X3 0.1104606 0.0006731273 164.10066        0 // ("Words")
X4 0.4397685 0.0006947700 632.96998        0 // ("LnTweetCount")
X5 0.0402905 0.0006583052  61.20337        0 // ("LnPopulation")
X6 0.3279818 0.0006649358 493.25341        0 // ("DayAge")

```

???
OK, so last minute I imported the data and started poking at R's utilities.  I didn't get any good plots out of it, as I think the confidence intervals might have been too small to appear on the charts.

But I did get these T-statistics and p values (2 tailed), which actually look pretty good!  We can see again that LnPop is not as good as the rest, though still looking okay.

Test is against null hypothesis, B_i = 0, so the numbers aren't confirmed, but we have a good idea there is indeed some linear relations between these properties.

---
layout:true
class: center, middle
---
![Words Vs Followers](fnl_words.png)
???
Okay, let's look at some pictures.

Jittering might be a good idea for words/links/population

Regression line for each plot is holding all other vars at their mean (zero).

---
![Links Vs Followers](fnl_links.png)

???

Nothing too interesting with the word count or link count regressions.

---
![LnPopulation Vs Followers](fnl_pop.png)
???
Population is pretty discrete: only 100 cities included.  1000 avail, but take a long time (distance calc).

I think the _lack_ of a relation here is actually interesting.

---
![Age Vs Followers](fnl_age.png)
???
The age graph has one of the most interesting anomalies: the dense cloud right around 1stD age.  Who are they?

---
![Tweets Vs Followers](fnl_tweets.png)
???
Here it almost looks like there's three separate groups, the cloud and the two spikes.

Maybe some color-coding of these dots for other vars would help determine if the anomalies had some other pattern

---
layout: false
class: center, middle

## Further Analysis

.topbox[
* More robust use of R's calculating and plotting tools
* Calculation/plotting of confidence intervals for regression coefficients
* Deeper examination of hypothesis testing, discussion of effect of variable inclusions on errors
* Use content analysis to provide more properties to analyze
* Try more varied property sets to find more interesting relations
* Explore varied plotting approaches
]
???
These are all things I basically half-did, but couldn't complete in time

R: Reinventing the wheel sounded cooler before I saw how awesome R's wheels are.  I have since added the ability to format my data to load into R.

Confidence intervals: I spent 7 hours going through Stats 414-415 trying to get this concept down.  I got it for simple LR, but hit a wall at Stat501 

se(y) = sqrt(MSE(Xht (Xt X)-1 Xh)).  Other lectures: "Use R to calculate this se(y).... -_-


HEY!  Since then, we've actually starting using R a bit, and gotten some p values!  Might leverage them into confidence intervals

Content analysis: wordnet is cool!  Noun/verb identification is hard

I saw a bunch of cool charts in plotting tools/walkthroughs, but didn't really have time to grok them sufficietly.

---
class: center, middle
# Q & A
.topbox[
* Eric Sorell
<br>
* Project source code available at https://github.com/sore0159/data_analysis
]

???
Even the source for this presentation itself!

