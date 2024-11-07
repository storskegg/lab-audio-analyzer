//
// Created by Liam Conrad on 11/6/24.
//

#include "binning.h"

double g(boundaries_t* b) {
  return log(b->f_max) - log(b->f_min);
}

double bin_frequency(int idx, boundaries_t* b, int num_bins) {
  return min * exp((double(idx) * g(b->f_min, b->f_max))/double(num_bins-1))
}

double to_precision(double n, int places) {
  if (places == 0) {
    return n;
  }
  float p = pow(10, places);
  return round(n * p) / p;
}

double* generate_bins(boundaries_t* b, int precision, size_t num_bins) {
  // At some point, disallow even number of bins
  if (num_bins % 2 == 0) {
    return NULL;
  }

  double* v = malloc(num_bins);

  for (size_t i = 0; i < num_bins; i++) {
    v[i] = to_precision(bin_frequency(i, b, num_bins), precision);
  }

  return v;
}
