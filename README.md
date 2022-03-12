# go-learn

### Simple example of machine learning in go using [go-learn package](https://github.com/sjwhitworth/golearn)

</br>

## Understanding matrix

### Vector

- ordered collection of numbers arranged in row or column
- Each number in vector is called component
- In go vectors are represented as slice-like defined type
- [Go num package](https://www.gonum.org/) provides easy way to use numerical types

### Matrices

- Rectangular organizations of numbers
- Linear algebra dictates the rules of manipulation with matrices
- Location in matrix is defined by `A`<sub>`row col`</sub> for rows and columns

<br>

<img src="./img/matrix.png" style="max-width:500px"/>

<br>

### Determinant

<br>

<img src="./img/determinant.png" style="max-width:500px"/>

<br>

## Evaluating models

- We can evaluate performace of our models by measuring how far off we were from predicted value

### Result types

- Continuos => result is probability, data for computation consists of numbers (stock price, temp, total sales)
- Categorical => Result is class that got most probability from dataset (Fraud / not fraud, name)

### Error values

#### MSE (Mean Squared Error)

- Takes the squares of the errors of correct value
- Tells us how different is in average prediction from predicted value

#### MAE (Mean Absolute Error)

- Maintains same units
- It is the difference between predicted value and actual value

#### MSD (Mean Squared Deviation)

### Evaluation metrics

- First you need to determine metric that fits the problem the best

#### Possible evaluation scenarios:

- True Positive
  - Predicted and observed category are same
  - Example: predicted fraud and observation was fraud
- False Positive
  - Predicted and observed category are not same
  - Example: predicted fraud but observation was not fraud
- True Negative
  - Predicted certain category, observation was not that category
  - Example: predicted not fraud and observation was not fraud
- False Negative
  - Predicted certain category, observation was the different category
  - Example predicted not fraud but observation was frau

#### Common metrics for measuring:

- Accuracy - Percentage of predictions that were right
- Precision - Percentage of positive predictions
- Recall - Percentage of identified positive predictions

### Individual numerical metrics

- Provide complete representation of model performance

#### Confusion matrices

- Use confusion matrices to determine the best algorithm
- Cols = Actual values
- Rows = Predicted values <br>
  +-----------------+-----------------+ <br>
  |&nbsp;&nbsp; True positives &nbsp;&nbsp; |&nbsp;&nbsp;False positives &nbsp;&nbsp;| <br>
  |&nbsp;&nbsp; False negatives |&nbsp;&nbsp;True Negatives &nbsp;| <br>
  +-----------------+-----------------+ <br>

<img src="./img/confusion-matrix.png" style="max-width:500px"/>

<br>

#### ROC (Receiver operating characteristic)

- Tells us overall picture of performance of binary classifiers
- Plot the recall versus false positive rate
- Thresholds represent various boundaries or rankings

<br>

<img src="./img/roc.png" style="max-width:500px"/>

<br>

- The model evaluated by the ROC curve makes a prediction for two classes based on probability, ranking or score.

<br>

<img src="./img/roc-curve.png" style="max-width:500px"/>

<br>

- A good ROC curve:
  - is in the upper left section of the plot, that means that model has better then random predictive power.
  - has more AUC (Area under curve) => if AUC is less then 50% predictions are to be considered random

<br>

## Linear Regression

<br>

<img src="./img/lin-reg.png" style="max-width:500px"/>

<br>

## Logistic Regression

## Understanding clustering

## Time series data

## Neural net

## Dockerizing model

## Sources

- [Confusion matrix](https://www.youtube.com/watch?v=Kdsp6soqA7o)
- [Machine learning with go](https://www.youtube.com/playlist?list=PLTgRMOcmRb3MgR1S-5DdMJyT6NzR_-7wE)
- [Simple Linear Regression for Machine Learning](https://www.youtube.com/watch?v=HoqXask9cN8)
