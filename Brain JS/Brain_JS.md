# Brain JS (Scrimba course)

## Training 

Use ```train()``` to train the network with an array of training data. The network has to be trained with all the data in bulk in one call to ```train()```. More training patterns will probably take longer to train, but will usually result in a network better at classifying new patterns.

### For training with NeuralNetwork

Each training pattern should have an ```input``` and an ```output```, both of which can be either an array of numbers from ```0``` to ```1``` or a hash of numbers from ```0``` to ```1```.

## How they learn - Structure

```JS
// inputs & outputs
(inputs) => outputs; // receive inputs as an arguments and produce outputs

// random values
Math.random(); // begins with the banch of random outputs

// The output is a random at first. It somewhere between 0 and 1

// activation "relu"
function relu(value) {
  return value < 0 ? 0 : value;
}

// In artificial neural networks, the activation function of a node defines the output of that node, or "neuron," given an input or set of inputs. This output is then used as input for the next node and so on until a desired solution to the original problem is found.

```

## How they learn - Layers

<img src = "https://img.scoop.it/rHvaJDfwTm-zNgvWAB8fcjl72eJkfbmt4t8yenImKBVvK0kTmF0xjctABnaLJIm9" width = 300px>

*Input* and *output* layers are configured **automatically**. 

Hidden Layers are configured by **programmer**.

In the image the input layer has four neurons, the first hidden layer - five neurons, the second hidden layer - three neurons, the output has one neuron.