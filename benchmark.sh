#!/bin/bash

ab -n 10000 -c 4 'http://127.0.0.1:3001/?test=payload'
