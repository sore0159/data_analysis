### Data Analysis Project
Dec 2016

This repository is an assemblage of the various bits of programming put together to perform and present data analysis for a job interview.

##Apps Directory:
* [Crunch](./apps/crunch/)
    The main application, organizing the usage of all the other libraries.  Used during the project to evaluate how well regressions were fitting the data, and generate png files of the various plots.

* [Present](./apps/present/)
    A slideshow app to present the project.  Basically a local http server that serves markdown-based slideshow using remark.js

* [Presentation Content](./apps/present/FILES/content.md)
        The slides, in markdown form.  Slides broken up by lines, with "???" denoting the start of presenters notes for each slide.

* [Tweet Stream](./apps/tweet_stream/)
    The application that monitered and recorded the tweets used for the data set


##Libraries:
* [Cities](./cities/)
    Loading data on US Cities, calculating great-circle distance to geographic points

* [Plot](./plot/)
    Different plots, mainly scatterplotting but also histograms, using the gonum plotter library

* [Maths](./maths/)
    Structuring the variables, utilities for statstics calculations and regressions.

* [Twitter](./twitter/)
    Library for the structuring of data from tweets, storing and loading to CSV files

* [Wordnet](./wordnet/)
    Unused start of content analysis, using the Princeton Wordnet database
