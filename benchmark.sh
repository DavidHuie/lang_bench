#!/bin/bash

ab -n 10000 -c 10 'http://127.0.0.1:3001/?test=payload'
