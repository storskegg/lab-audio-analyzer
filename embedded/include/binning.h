//
// Created by Liam Conrad on 11/6/24.
//

#ifndef EMBEDDED_BINNING_H
#define EMBEDDED_BINNING_H

#include <math.h>
#include <stdlib.h>
#include <vector>

double g(double min, double max);
double binFrequency(int idx, double min, double max, int num_bins);

double to_precision(double n, int places);

double* generate_bins(boundaries_t* boundaries, size_t num_bins);

typedef struct boundaries_t {
  double f_min;
  double f_max;
} boundaries_t;

#endif //EMBEDDED_BINNING_H
