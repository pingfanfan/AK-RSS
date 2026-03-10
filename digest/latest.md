# OPMLWatch Digest

Generated at: 2026-03-10 17:14:35 UTC

- [Writing an LLM from scratch, part 28 -- training a base model from scratch on an RTX 3090](https://www.gilesthomas.com/2025/12/llm-from-scratch-28-training-a-base-model-from-scratch) (`gilesthomas.com`)
  - TLDR: 个人开发者用单张RTX 3090成功训练163M参数GPT-2级模型，使用FineWeb数据集，验证了低成本训练基础模型的可行性。
  - WHAT: 训练了一个163M参数的GPT-2小型基础模型，使用Hugging Face FineWeb数据集，在RTX 3090上完成完整训练，验证损失值接近原始GPT-2论文报告。
  - WHY: 受Andrej Karpathy的nanochat项目启发，结合Sebastian Raschka书籍方法，验证消费级GPU训练基础模型的实际可行性，挑战‘必须依赖大集群’的假设。
  - ACTION: 使用RTX 3090或同级别GPU，下载FineWeb数据集，参考开源代码（如nanochat或本系列），尝试训练小型模型（<200M参数），监控损失曲线并调整超参数。
  - TWEET: 个人开发者福音：RTX 3090训练163M参数LLM成功！成本可控，性能接近GPT-2。技术细节与数据集见链接。
