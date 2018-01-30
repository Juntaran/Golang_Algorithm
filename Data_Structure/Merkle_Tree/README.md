# Merkle Tree

## Merkle Tree 特性

Merkle Tree 又称 Merkle Hash Tree  

1. 本质为二叉树，叶子节点存储数据  
2. 非叶子节点根据它的两个孩子的值进行 hash 计算  
3. 整棵树的根结点存储整个 Tree 的最终 Hash 值  

## Merkle Tree 与区块链

区块链的区块利用 Merkle Tree 把区块重的大量信息缩成一个 Hash 值  
这个值会被记录到区块头中被验证  
