//
// Created by Liam Conrad on 11/6/24.
//

#include "binning.h"

double g(double min, double max) {
  return log(max) - log(min);
}

double binFrequency(int idx, double min, double max, int numBins) {
  return min * exp((double(idx) * g(min, max))/double(numBins-1))
}
