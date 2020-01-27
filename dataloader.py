"""
    Dataloader Class
"""

import torchvision.datasets as ds
import numpy as np
import torch

class MNIST_Dataset():
    def __init__(self, test_val_split=0.5, normalize=True):
        """
        arg:
            test_val_split: the proportion of data_test is used for test data
        """
        self.idx = 0
        self.ds = ds
        self.ratio = test_val_split
        
        self._extract_train_test_data()
        self._test_val_split()
        if normalize:
            self._normalize()
    
    def _extract_train_test_data(self):
        data_train = self.ds.MNIST(root='/Users/ellise/Desktop/datasets/MNIST_data/', train=True, download=False)
        data_test = self.ds.MNIST(root='/Users/ellise/Desktop/datasets/MNIST_data/', train=False, download=False)
        self.X_train, self.y_train = data_train.train_data.float(), data_train.train_labels
        self.X_test, self.y_test = data_test.test_data.float(), data_test.test_labels
        self.num_instances_test = self.X_test.shape[0]
        
    def _test_val_split(self):
        """
        split test and validation set according to split ratio and total_size
        """
        split_idx = self.num_instances_test - int(self.num_instances_test * self.ratio)
        self.X_val, self.y_val = self.X_test[:split_idx], self.y_test[:split_idx]
        self.X_test, self.y_test = self.X_test[split_idx:], self.y_test[split_idx:]
        print("Train set size: {}; Test set size: {}; Validation set size: {}".format(self.X_train.shape[0],\
                                                                                     self.X_test.shape[0],\
                                                                                     self.X_val.shape[0]))
        
    def get_train_batch(self, batch_size):
        idx_start = self.idx
        idx_end = self.idx + batch_size
        self.idx += batch_size

        if(self.idx >= self.X_train.shape[0]):
            self.idx = 0
            idx_list = np.linspace(0,self.X_train.shape[0]-1,self.X_train.shape[0])
            idx_list_shuffle = np.random.shuffle(np.linspace(0,self.X_train.shape[0]-1,self.X_train.shape[0]))

            self.X_train[idx_list] = self.X_train[idx_list_shuffle]
            self.y_train[idx_list] = self.y_train[idx_list_shuffle]
        return (self.X_train[idx_start:idx_end], self.y_train[idx_start:idx_end])
    
    def _normalize(self):
        def squash(X):
            # put data into 0-1
            return (X - torch.min(X)) / (torch.max(X) - torch.min(X))
        self.X_train = squash(self.X_train)
        print()
        self.X_test = squash(self.X_test)
        self.X_val = squash(self.X_val)
        
    def get_test_data(self):
        return (self.X_test, self.y_test)
        
    def get_validation_data(self):
        return (self.X_val, self.y_val)
    
    def get_num_of_instances_in_train(self):
        # used for return how many instances in training set in total
        return self.X_train.shape[0]
    
    def get_test_set_size(self):
        return self.X_test.shape[0]
    
    def get_val_set_size(self):
        return self.X_val.shape[0]