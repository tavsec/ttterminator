# TTTerminator

This repository store the code for TTTerminator, an service which 
can play the [TTTM](https://tttm.dev) challenge.

## About

This bot is implemented using [Alpha-Beta Pruning](https://en.wikipedia.org/wiki/Alpha%E2%80%93beta_pruning) algorithm.

## Usage
```bash
curl http://localhost:8080/move?\
gid=ff66ae0f-3c35-418c-9634-0aacaaecd523&\
size=3&\
playing=O&\
moves=X-1-1_O-0-0_X-0-2_O-2-0_X-2-2
```