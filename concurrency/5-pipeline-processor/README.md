# Pipeline Processor
## Pipeline Processor

Build a multi-stage concurrent pipeline with fan-out and fan-in patterns.

### Recommended Prerequisites

Complete these exercises first for the best learning experience:

- Worker Pool

## Background

**Pipeline Pattern**: Chain processing stages where output of one stage feeds into the next, connected by channels.

**Fan-Out**: Multiple goroutines read from the same channel to parallelize work.

**Fan-In**: Multiple channels are merged into one output channel.

Pipelines enable concurrent data processing with clear separation of concerns and efficient resource usage.

## Your Task

Build a 3-stage data processing pipeline:

1. **Stage 1 (Parse)**: Parse job data
2. **Stage 2 (Transform)**: Transform/process the data (fan-out to 2 workers)
3. **Stage 3 (Validate)**: Validate results
4. **Merge**: Combine results from all stage 3 workers

## Key Concepts

- **Pipeline Pattern**: Chain stages with channels for concurrent data flow
- **Fan-Out**: Multiple workers read from same channel to parallelize work
- **Fan-In (Merge)**: Combine multiple channels into one output channel
- **Channel Ownership**: Each stage owns its output channel and closes it when done
- **Composition**: Build complex workflows by chaining simple stages

