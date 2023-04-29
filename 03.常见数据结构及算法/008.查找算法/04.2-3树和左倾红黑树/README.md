# 2-3树和左倾红黑树
红黑树是一种近似平衡的二叉查找树,从`2-3`树或`2-3-4`树衍生而来.通过对二叉树节点进行染色,染色为红或黑节点,来模仿`2-3`树或`2-3-4`树的3节点和4节点,从而让树的高度减少.`2-3-4`树对照现实的红黑树是普通的红黑树,而`2-3`树对照实现的红黑树是一种变种,称为左倾红黑树,其更容易实现.

使用平衡树数据结构,可以提高查找元素的速度.

# `2-3`树
## 介绍
`2-3`树是一棵严格自平衡的多路查找树,又称`3阶的B树`

它不是一棵二叉树,是一棵三叉树,具有以下特征:
1. 内部节点要么有1个数据元素和2个孩子,要么有2个数据元素和3个孩子,但有1或2个数据元素.
2. 所有叶子节点到根节点的长度一致,这个特征保证了完全平衡,非常完美的平衡.
3. 每个节点的数据元素保持从小到大的排序,两个数据元素之间的子树的所有值大小介于两个数据元素之间.

因为`2-3`树的第二个特征,它是一个完美平衡的树,非常完美,除了叶子节点,其他的节点都没有空儿子,所以树的高度非常的小.如图:
![2_3_tree.png](./assets/2_3_tree.png)

如果一个内部节点拥有一个数据元素,两个子节点,则此节点为2节点.如果一个内部节点拥有2个数据元素,3个节点,则此节点为3节点.

可以说,所有平衡树的核心都在于插入和删除逻辑,我们主要分析这两个操作.

## `2-3`树插入元素
在插入元素时,需要先找到插入的位置,使用二分查找从上自下查找树节点.

找到插入位置时,将元素插入该位置,然后进行调整,使得满足`2-3`树的特征,主要有3种情况:
1. 插入元素到一个2节点,直接插入即可,这样节点变成3节点.
2. 插入元素到一个3节点,该3节点的父亲是一个2节点,先将节点变成临时的4节点,然后向上分裂调整一次.
3. 插入元素到一个3节点,该3节点的父亲是一个3节点,先将节点变成临时的4节点,然后向上分裂调整,此时父亲节点变为临时4节点,继续向上分裂调整.

如图(来自维基百科):
![2-3_insertion.png](./assets/2-3_insertion.png)

核心在于插入3节点后,该节点变为临时4节点,然后进行分裂恢复树的特征.最坏情况为插入节点后,每一次分裂后都导致上一层变为临时4节点,直到树根节点,这样需要不断向上分裂.

临时4节点的分裂,细分为6种情况,如图:
![2-3_tree_split.png](./assets/2-3_tree_split.png)

与其他二叉查找树由上而下生长不同,`2-3`树是从下至上的生长

`2-3`树的实现将会放在`B树`章节

## `2-3`树删除元素
删除操作就复杂得多.

`2-3`树的特征注定它是一棵非常完美平衡的三叉树,其所有子树也都是完美平衡,所以`2-3`树的节点的儿子,要么都是空儿子,要么都不是空儿子,比如`2-3`树的某个节点`A`有两个儿子`B`和`C`,儿子`B,C`要么都没有孩子,要么孩子都是满的,不然`2-3`树所有叶子节点到根节点的长度一致这个特征就被破坏了.

基于上面的现实,我们来分析删除的不同情况,删除中间节点和叶子节点.

情况1:删除中间节点
> 删除的非叶子节点,该节点一定是有2棵或3棵子树的,那么从子树中找到其最小后继节点,该节点是叶子节点,用该节点替换被删除的非叶子节点,然后再删除这个叶子节点,进入情况2.

> 如果还找到最小后继节点,当有2棵子树时,那么从右子树一直往左下方找,如果有3棵子树,被删除节点在左边,那么从中子树一直往左下方找,否则从右子树一直往下方找.

情况2:删除叶子节点
> 删除的叶子节点,这时如果叶子节点是3节点,那么直接变为2节点即可,不影响平衡,但是,如果叶子节点是2节点,那么删除后,其父亲节点将会缺失一个儿子,破坏了满孩子的`2-3`树特征,需要进行调整后才能删除.

> 针对情况2,删除一个2节点叶子节点,会导致父亲节点缺失一个儿子,破坏了`2-3`树的特征,我们可以进行调整变换,主要有两种调整:
1. 重新分布:尝试从兄弟节点哪里借值,然后重新调整节点.
2. 合并:如果兄弟借不到值,合并节点,再向上递归处理

如图:
![2-3_tree_delete1.jpg](./assets/2-3_tree_delete1.jpg)

如果被删除的叶子节点有兄弟是3节点,那么从兄弟那里借一个值填补被删除的叶子节点,然后兄弟和父亲重新分布调整位置,下面是重新分布的具体例子:
![2-3_tree_delete2.jpg](./assets/2-3_tree_delete2.jpg)

可以看到,删除`100`,从兄弟那里借来一个值`80`,然后重新调整父亲,兄弟们的位置.

如果兄弟们都是2节点呢?那么就合并节点:将父亲和兄弟节点合并,如果父亲是2节点,那么父亲就留空了,否则父亲就从3节点变成2节点,下面是合并的两个具体例子:
![2-3_tree_delete3.jpg](./assets/2-3_tree_delete3.jpg)

可以看到,删除`80`,而兄弟节点`60`和父亲节点`90`都是2节点,所以父亲下来和兄弟合并,然后父亲节点变为空节点.
![2-3_tree_delete6.jpg](./assets/2-3_tree_delete6.jpg)

可以看到,删除`70`而兄弟节点都为2节点,父亲节点为3节点,那么父亲下来和其中一个兄弟合并,然后父亲从3节点变为2节点.

但是,如果合并后,父亲节点变空了,也就是说中间节点留空要怎么办,那么可以继续递归处理,如图:
![2-3_tree_delete4.jpg](./assets/2-3_tree_delete4.jpg)

中间节点是空的,那么可以继续从兄弟那里借节点或者和父亲合并,直到根节点,如果到达了根节点呢.如图:
![2-3_tree_delete5.jpg](./assets/2-3_tree_delete5.jpg)

递归到了根节点后,如果存在空的根节点,我们可以直接把该空节点删除即可,这时树的高度减少一层.

# 左倾红黑树
## 介绍
左倾红黑树可以由`2-3`树的二叉树形式来实现

其定义为:
1. 根节点的链接是黑色的
2. 红链接均为左链接
3. 没有任何一个节点同时和两条红链接相连
4. 任意一个节点到达叶子节点的所有路径,经过的黑链接数量相同,也就是该树是完美黑色平衡的.比如,某一个节点,它可以到达5个叶子节点,那么这5调路径上的黑链接数量一样.

由于红链接都在左边,所以这种红黑树又称左倾红黑树.左倾红黑树与`2-3`树一一对应,只要将左倾链接画平,如图:
![llrb_tree1.jpg](./assets/llrb_tree1.jpg)

## 节点旋转和颜色转换
首先,定义树的结构`LLRBTree`,以及表示左倾红黑树的节点`LLRBTNode`:
```go
// 定义颜色
const (
    RED   = true
    BLACK = false
)
// 左倾红黑树
type LLRBTree struct {
    Root *LLRBTNode // 树根节点
}
// 左倾红黑树节点
type LLRBTNode struct {
    Value       int64     // 值
    Times       int64      // 值出现的次数
    Left        *LLRBTNode // 左子树
    Right       *LLRBTNode // 右子树
    Color       bool       // 父亲指向该节点的链接颜色
}
// 新建一棵空树
func NewLLRBTree() *LLRBTree {
    return &LLRBTree{}
}
// 节点的颜色
func IsRed(node *LLRBTNode) bool {
    if node == nil {
        return false
    }
    return node.Color == RED
}
```

在节点`LLRBTNode`中,我们存储的元素字段为`Value`,由于可能由重复的元素插入,所以多了一个`Times`字段,表示该元素出现几次.

当然,红黑树中的红黑颜色使用`Color`定义,表示父亲指向该节点的链接颜色,为了方便,我们还构造了一个辅助函数`IsRed()`

在元素添加和实现的过程中,需要做调整操作,有两种旋转操作,对某节点的右链接进行左旋转,或者左链接进行右旋转.

如图是对节点`h`的右链接进行左旋转:
![llrb_tree_left_rotate.jpg](./assets/llrb_tree_left_rotate.jpg)

代码实现:
```go
// 左旋转
func RotateLeft(h *LLRBTNode) *LLRBTNode {
    if h == nil {
        return nil
    }
    // 看图理解
    x := h.Right
    h.Right = x.Left
    x.Left = h
    x.Color = h.Color
    h.Color = RED
    return x
}
```

如图是对节点`h`的左链接进行右旋转:
![llrb_tree_right_rotate.jpg](./assets/llrb_tree_right_rotate.jpg)

代码实现:
```go
// 右旋转
func RotateRight(h *LLRBTNode) *LLRBTNode {
    if h == nil {
        return nil
    }
    // 看图理解
    x := h.Left
    h.Left = x.Right
    x.Right = h
    x.Color = h.Color
    h.Color = RED
    return x
}
```

由于左倾红黑树不允许一个节点有两个红链接,所以需要做颜色转换,如图:
![llrb_tree_color_change.jpg](./assets/llrb_tree_color_change.jpg)

代码实现:
```go
// 颜色转换
func ColorChange(h *LLRBTNode) {
    if h == nil {
        return
    }
    h.Color = !h.Color
    h.Left.Color = !h.Left.Color
    h.Right.Color = !h.Right.Color
}
```

## 添加元素实现
每次添加元素节点时,都将该节点`color`字段,也就是父亲指向它的链接设置为`RED`红色

接着判断其父亲是否有两个红链接(如果连续的两个左红链接或者左右红色链接),或者有右红色链接,进行颜色变换或旋转操作.

主要有以下这几种情况:

插入元素到2节点,直接让节点变为3节点,不过当右插入时需要左旋转使得红色链接在左边,如图:
![llrb_tree_insert_2node.jpg](./assets/llrb_tree_insert_2node.jpg)

插入元素到3节点,需要做旋转和颜色转换操作,如图:
![llrb_tree_insert_3node.jpg](./assets/llrb_tree_insert_3node.jpg)

也就是说,在一个已经是红色左链接的节点,插入一个新节点的状态变化如下:
![llrb_tree_red_change.jpg](./assets/llrb_tree_red_change.jpg)

根据上述的演示以及旋转,颜色转换等操作,添加元素的代码为:
```go
// 左倾红黑树添加元素
func (tree *LLRBTree) Add(value int64) {
    // 跟节点开始添加元素，因为可能调整，所以需要将返回的节点赋值回根节点
    tree.Root = tree.Root.Add(value)
    // 根节点的链接永远都是黑色的
    tree.Root.Color = BLACK
}
// 往节点添加元素
func (node *LLRBTNode) Add(value int64) *LLRBTNode {
    // 插入的节点为空，将其链接颜色设置为红色，并返回
    if node == nil {
        return &LLRBTNode{
            Value: value,
            Color: RED,
        }
    }
    // 插入的元素重复
    if value == node.Value {
        node.Times = node.Times + 1
    } else if value > node.Value {
        // 插入的元素比节点值大，往右子树插入
        node.Right = node.Right.Add(value)
    } else {
        // 插入的元素比节点值小，往左子树插入
        node.Left = node.Left.Add(value)
    }
    // 辅助变量
    nowNode := node
    // 右链接为红色，那么进行左旋，确保树是左倾的
    // 这里做完操作后就可以结束了，因为插入操作，新插入的右红链接左旋后，nowNode节点不会出现连续两个红左链接，因为它只有一个左红链接
    if IsRed(nowNode.Right) && !IsRed(nowNode.Left) {
        nowNode = RotateLeft(nowNode)
    } else {
        // 连续两个左链接为红色，那么进行右旋
        if IsRed(nowNode.Left) && IsRed(nowNode.Left.Left) {
            nowNode = RotateRight(nowNode)
        }
        // 旋转后，可能左右链接都为红色，需要变色
        if IsRed(nowNode.Left) && IsRed(nowNode.Right) {
            ColorChange(nowNode)
        }
    }
    return nowNode
}
```

## 添加元素算法分析
参考论文:https://www.cs.princeton.edu/~rs/talks/LLRB/LLRB.pdf

左倾红黑树的最坏树高度为`2log(n)`,其中`n`为树的节点数量.为什么呢?我们先把左倾红黑树当作`2-3`树,也就是说最坏情况下沿着`2-3`树左边的节点都是3节点,其他节点都是2节点,这时树高近似`log(n)`,再从`2-3`树转成左倾红黑树,当3节点不画平时,可以知道树高度变成原来`2-3`树树高的2倍.虽然如此,构造一棵最坏的左倾红黑树很难.

`AVL`树的最坏树高度为`1.44log(n)`.由于左倾红黑树是近似平衡的二叉树,没有`AVL`树的严格平衡,树的高度会更高一点,因此查找操作效率比`AVL`树低,但时间复杂度只在于常数项的差别,去掉常数项,时间复杂度仍然是`log(n)`.

我们的代码实现中,左倾红黑树的插入,需要逐层判断是否需要旋转和变色,复杂度为`log(n)`,当旋转变色后导致上层存在连续的红左链接或者红色左右链接,那么需要继续旋转和变色,可能有多次这种调整操作,如图在箭头处添加新节点,出现了右红链接,要一直向上变色到根节点.
![llrb_tree_example.jpg](./assets/llrb_tree_example.jpg)

我们可以优化代码,使得在某一层旋转变色后,如果其父层没有连续的左红链接或者不需要变色,那么可以直接退出,不需要逐层判断是否需要旋转和变色.

对于`AVL`树来说,插入最多旋转2次,但其需要逐层更新树高度,复杂度也是为`log(n)`

按照插入效率来说,很多教程都说左倾红黑树会比`AVL`树好一点,因为其不要求严格的平衡,会插入得更快点,但根据我们实际上的递归代码,2者都需要逐层向上判断是否需要调整,只不过`AVL`树多了更新树高度的操作,此操作影响了一点点效率,但我觉得2种树的插入效率都差不多.

## 删除元素实现
1. 情况:如果删除的非叶子节点,找到其最小后驱节点,也就是在其右子树种一直向左找,找到该叶子节点替换被删除的节点,然后删除该叶子节点,变成情况2
2. 情况:如果删除的是叶子节点,如果它是红节点,也就是父亲指向它的链接为红色,那么直接删除即可,否则我们需要进行调整,使它变为红色,再删除.

在这里,为了使得删除叶子节点时可以直接删除,叶子节点必须变为红节点.

我们创造2种操作,如果删除的节点在左子树中,可能需要进行红色左移,如果删除的节点在右子树中,可能需要进行红色右移.

介绍红色左移的步骤:
- 要在树`h`的左子树中删除元素,这时树`h`根节点是红色,其儿子`b,d`节点都为黑色节点,且两个黑色节点都是2节点,都没有左红孩子,那么直接对`h`树根节点变色即可,如图:
![llrb_tree_delete1.jpg](./assets/llrb_tree_delete1.jpg)

如果存在右儿子`d`是3节点,有左红孩子`e`,那么需要先对`h`树根节点变色后,对右儿子`d`右旋,再对`h`树根节点左旋,最后再一次对`h`树根节点变色,如图:
![llrb_tree_delete2.jpg](./assets/llrb_tree_delete2.jpg)

红色左移可以总结为下图:
![llrb_tree_delete3.jpg](./assets/llrb_tree_delete3.jpg)

代码如下:
```go
// 红色左移
// 节点 h 是红节点，其左儿子和左儿子的左儿子都为黑节点，左移后使得其左儿子或左儿子的左儿子有一个是红色节点
func MoveRedLeft(h *LLRBTNode) *LLRBTNode {
    // 应该确保 isRed(h) && !isRed(h.left) && !isRed(h.left.left)
    ColorChange(h)
    // 右儿子有左红链接
    if IsRed(h.Right.Left) {
        // 对右儿子右旋
        h.Right = RotateRight(h.Right)
        // 再左旋
        h = RotateLeft(h)
        ColorChange(h)
    }
    return h
}
```

为什么要红色左移,是要保证调整后,子树根节点`h`的左儿子或左儿子的左儿子有一个红色节点,这样从`h`的左子树递归删除元素蔡可以继续下去.

红色右移的步骤类似,如图:
![llbr_tree_delete4_2.jpg](./assets/llbr_tree_delete4_2.jpg)

代码如下:
```go
// 红色右移
// 节点 h 是红节点，其右儿子和右儿子的左儿子都为黑节点，右移后使得其右儿子或右儿子的右儿子有一个是红色节点
func MoveRedRight(h *LLRBTNode) *LLRBTNode {
    // 应该确保 isRed(h) && !isRed(h.right) && !isRed(h.right.left);
    ColorChange(h)
    // 左儿子有左红链接
    if IsRed(h.Left.Left) {
        // 右旋
        h = RotateRight(h)
        // 变色
        ColorChange(h)
    }
    return h
}
```

为什么要红色右移,同样是为了保证树根节点`h`的右儿子或右儿子的右儿子有一个是红色节点,往右子树递归删除元素可以继续下去.

明确以下到底是如何删除元素的.

我们知道`2-3`树的删除是从叶子节点开始,自底向上的向兄弟节点借值,或和父亲合并,然后一直递归到根节点.左倾红黑树参考了这种做法,但更巧妙,左倾红黑树要保证一路上每次递归进入删除操作的子树树根一定是一个3节点,所以需要适当的红色左移或右移,这样一直递归到叶子节点,叶子节点也会是一个3节点,然后就可以直接删除叶子节点,最后再自底向上的恢复左倾红黑树的特征.

下面是左倾红黑树从`树h`删除元素的示例图,往`树h`左子树和右子树删除元素分别有4种情况,后2种情况需要使用到红色左移或右移,状态演变之后,`树h`才可以从左或右子树进入下一次递归:
![llrb_tree_delete_total.jpg](./assets/llrb_tree_delete_total.jpg)

可以对照大图,继续阅读下面左倾红黑树删除元素代码:
```go
// 左倾红黑树删除元素
func (tree *LLRBTree) Delete(value int64) {
    // 当找不到值时直接返回
    if tree.Find(value) == nil {
        return
    }
    if !IsRed(tree.Root.Left) && !IsRed(tree.Root.Right) {
        // 左右子树都是黑节点，那么先将根节点变为红节点，方便后面的红色左移或右移
        tree.Root.Color = RED
    }
    tree.Root = tree.Root.Delete(value)
    // 最后，如果根节点非空，永远都要为黑节点，赋值黑色
    if tree.Root != nil {
        tree.Root.Color = BLACK
    }
}
```

首先`tree.Find(value)`找到可以删除的值时才能进行删除.

当根节点的左右子树都为黑色节点时,那么先将根节点变为红节点,方便后面的红色左移或右移.

删除完节点:`tree.Root=tree,Delete(value)`后,需要将根节点染回黑色,因为左倾红黑树的特征之一是根节点永远都是黑色.

核心的从子树中删除元素代码如下:
```go
// 对该节点所在的子树删除元素
func (node *LLRBTNode) Delete(value int64) *LLRBTNode {
    // 辅助变量
    nowNode := node
    // 删除的元素比子树根节点小，需要从左子树删除
    if value < nowNode.Value {
        // 因为从左子树删除，所以要判断是否需要红色左移
        if !IsRed(nowNode.Left) && !IsRed(nowNode.Left.Left) {
            // 左儿子和左儿子的左儿子都不是红色节点，那么没法递归下去，先红色左移
            nowNode = MoveRedLeft(nowNode)
        }
        // 现在可以从左子树中删除了
        nowNode.Left = nowNode.Left.Delete(value)
    } else {
        // 删除的元素等于或大于树根节点
        // 左节点为红色，那么需要右旋，方便后面可以红色右移
        if IsRed(nowNode.Left) {
            nowNode = RotateRight(nowNode)
        }
        // 值相等，且没有右孩子节点，那么该节点一定是要被删除的叶子节点，直接删除
        // 为什么呢，反证，它没有右儿子，但有左儿子，因为左倾红黑树的特征，那么左儿子一定是红色，但是前面的语句已经把红色左儿子右旋到右边，不应该出现右儿子为空。
        if value == nowNode.Value && nowNode.Right == nil {
            return nil
        }
        // 因为从右子树删除，所以要判断是否需要红色右移
        if !IsRed(nowNode.Right) && !IsRed(nowNode.Right.Left) {
            // 右儿子和右儿子的左儿子都不是红色节点，那么没法递归下去，先红色右移
            nowNode = MoveRedRight(nowNode)
        }
        // 删除的节点找到了，它是中间节点，需要用最小后驱节点来替换它，然后删除最小后驱节点
        if value == nowNode.Value {
            minNode := nowNode.Right.FindMinValue()
            nowNode.Value = minNode.Value
            nowNode.Times = minNode.Times
            // 删除其最小后驱节点
            nowNode.Right = nowNode.Right.DeleteMin()
        } else {
            // 删除的元素比子树根节点大，需要从右子树删除
            nowNode.Right = nowNode.Right.Delete(value)
        }
    }
    // 最后，删除叶子节点后，需要恢复左倾红黑树特征
    return nowNode.FixUp()
}
```

这段核心代码十分复杂,会用到红色左移和右移,当删除的元素小于根节点时,我们明白要在左子树中删除.如:
```go
    // 删除的元素比子树根节点小，需要从左子树删除
    if value < nowNode.Value {
        // 因为从左子树删除，所以要判断是否需要红色左移
        if !IsRed(nowNode.Left) && !IsRed(nowNode.Left.Left) {
            // 左儿子和左儿子的左儿子都不是红色节点，那么没法递归下去，先红色左移
            nowNode = MoveRedLeft(nowNode)
        }
        // 现在可以从左子树中删除了
        nowNode.Left = nowNode.Left.Delete(value)
    }
```

递归删除左子树前:`nowNode.Left=nowNode.Left.Delete(value)`,要确保删除的左子树根节点是红色节点,或左子树根节点的左儿子是红色节点,才能够继续递归下去,所以使用了`!IsRed(nowNode.Left) && !IsRed(nowNode.Left.Left)`来判断是否需要红色左移.

如果删除的值不小于根节点,那么进入以下逻辑:
```go
        // 删除的元素等于或大于树根节点
        // 左节点为红色，那么需要右旋，方便后面可以红色右移
        if IsRed(nowNode.Left) {
            nowNode = RotateRight(nowNode)
        }
        // 值相等，且没有右孩子节点，那么该节点一定是要被删除的叶子节点，直接删除
        // 为什么呢，反证，它没有右儿子，但有左儿子，因为左倾红黑树的特征，那么左儿子一定是红色，但是前面的语句已经把红色左儿子右旋到右边，不应该出现右儿子为空。
        if value == nowNode.Value && nowNode.Right == nil {
            return nil
        }
        // 因为从右子树删除，所以要判断是否需要红色右移
        if !IsRed(nowNode.Right) && !IsRed(nowNode.Right.Left) {
            // 右儿子和右儿子的左儿子都不是红色节点，那么没法递归下去，先红色右移
            nowNode = MoveRedRight(nowNode)
        }
        // 删除的节点找到了，它是中间节点，需要用最小后驱节点来替换它，然后删除最小后驱节点
        if value == nowNode.Value {
            minNode := nowNode.Right.FindMinValue()
            nowNode.Value = minNode.Value
            nowNode.Times = minNode.Times
            // 删除其最小后驱节点
            nowNode.Right = nowNode.Right.DeleteMin()
        } else {
            // 删除的元素比子树根节点大，需要从右子树删除
            nowNode.Right = nowNode.Right.Delete(value)
        }
```

首先,需要先判断该节点的左子根节点是否为红色节点`IsRed(nowNode.Left)`,如果是的话需要右旋:`nowNode = RotateRight(nowNode)`,将红色节点右旋是为了后面可以递归进入右子树.

然后,判断删除的值是否等于当前根节点的值,且其没有右节点:`value==nowNode.Value && nowNode.Right == nil`,如果是,那么该节点是要被删除的叶子节点,直接删除即可.

接着,判断是否需要红色右移:`IsRed(nowNode.Right) && !IsRed(nowNode.Right.Left)`,如果该节点右儿子和右儿子的左儿子都不是红色节点,那么没法递归进入右子树,需要红色右移,必须确保其右子树或右子树的左儿子有一个是红色节点.

再接着,需要判断是否找到了要删除的节点:`value==nowNode.Value`,找到时表示要删除的节点处于内部节点,需要用最小后驱节点来替换它,然后删除最小后驱节点.

找到最小后驱节点:`minNode:=nowNode.Right.FindMinValue()`之后,将最小后驱节点与要删除的内部节点替换,然后删除最小后驱节点:`nowNode.Right = nowNode.Right.DeleteMin()`,删除最小节点代码如下:
```go
// 对该节点所在的子树删除最小元素
func (node *LLRBTNode) DeleteMin() *LLRBTNode {
    // 辅助变量
    nowNode := node
    // 没有左子树，那么删除它自己
    if nowNode.Left == nil {
        return nil
    }
    // 判断是否需要红色左移，因为最小元素在左子树中
    if !IsRed(nowNode.Left) && !IsRed(nowNode.Left.Left) {
        nowNode = MoveRedLeft(nowNode)
    }
    // 递归从左子树删除
    nowNode.Left = nowNode.Left.DeleteMin()
    // 修复左倾红黑树特征
    return nowNode.FixUp()
}
```

因为最小节点在最左边的叶子节点,所以只需要适当的红色左移,然后一直左子树递归即可,递归完后需要修复左倾红黑树特征`nowNode.FixUp()`,代码如下:
```go
// 修复左倾红黑树特征
func (node *LLRBTNode) FixUp() *LLRBTNode {
    // 辅助变量
    nowNode := node
    // 红链接在右边，左旋恢复，让红链接只出现在左边
    if IsRed(nowNode.Right) {
        nowNode = RotateLeft(nowNode)
    }
    // 连续两个左链接为红色，那么进行右旋
    if IsRed(nowNode.Left) && IsRed(nowNode.Left.Left) {
        nowNode = RotateRight(nowNode)
    }
    // 旋转后，可能左右链接都为红色，需要变色
    if IsRed(nowNode.Left) && IsRed(nowNode.Right) {
        ColorChange(nowNode)
    }
    return nowNode
}
```

如果不是删除内部节点,依然是从右子树继续递归:
```go
    // 删除的元素比子树根节点大，需要从右子树删除
    nowNode.Right = nowNode.Right.Delete(value)
```

当然,递归完成后还要进行一次`FixUp()`,恢复左倾红黑树的特征.

## 删除元素算法分析
左倾红黑树删除元素需要自顶向下的递归,可能不断地红色左移和右移,也就是有很多的旋转,当删除的叶子节点后,还需要逐层恢复左倾红黑树的特征,时间复杂度仍然是和树高度有关:`log(h)`

## 查找元素等实现
查找最小值,最大值,或者某个值,代码如下:
```go
// 找出最小值的节点
func (tree *LLRBTree) FindMinValue() *LLRBTNode {
    if tree.Root == nil {
        // 如果是空树，返回空
        return nil
    }
    return tree.Root.FindMinValue()
}
func (node *LLRBTNode) FindMinValue() *LLRBTNode {
    // 左子树为空，表面已经是最左的节点了，该值就是最小值
    if node.Left == nil {
        return node
    }
    // 一直左子树递归
    return node.Left.FindMinValue()
}
// 找出最大值的节点
func (tree *LLRBTree) FindMaxValue() *LLRBTNode {
    if tree.Root == nil {
        // 如果是空树，返回空
        return nil
    }
    return tree.Root.FindMaxValue()
}
func (node *LLRBTNode) FindMaxValue() *LLRBTNode {
    // 右子树为空，表面已经是最右的节点了，该值就是最大值
    if node.Right == nil {
        return node
    }
    // 一直右子树递归
    return node.Right.FindMaxValue()
}
// 查找指定节点
func (tree *LLRBTree) Find(value int64) *LLRBTNode {
    if tree.Root == nil {
        // 如果是空树，返回空
        return nil
    }
    return tree.Root.Find(value)
}
func (node *LLRBTNode) Find(value int64) *LLRBTNode {
    if value == node.Value {
        // 如果该节点刚刚等于该值，那么返回该节点
        return node
    } else if value < node.Value {
        // 如果查找的值小于节点值，从节点的左子树开始找
        if node.Left == nil {
            // 左子树为空，表示找不到该值了，返回nil
            return nil
        }
        return node.Left.Find(value)
    } else {
        // 如果查找的值大于节点值，从节点的右子树开始找
        if node.Right == nil {
            // 右子树为空，表示找不到该值了，返回nil
            return nil
        }
        return node.Right.Find(value)
    }
}
// 中序遍历
func (tree *LLRBTree) MidOrder() {
    tree.Root.MidOrder()
}
func (node *LLRBTNode) MidOrder() {
    if node == nil {
        return
    }
    // 先打印左子树
    node.Left.MidOrder()
    // 按照次数打印根节点
    for i := 0; i <= int(node.Times); i++ {
        fmt.Println(node.Value)
    }
    // 打印右子树
    node.Right.MidOrder()
}
```

## 验证是否是一棵左倾红黑树
```go
// 验证是不是棵左倾红黑树
func (tree *LLRBTree) IsLLRBTree() bool {
    if tree == nil || tree.Root == nil {
        return true
    }
    // 判断树是否是一棵二分查找树
    if !tree.Root.IsBST() {
        return false
    }
    // 判断树是否遵循2-3树，也就是红链接只能在左边，不能连续有两个红链接
    if !tree.Root.Is23() {
        return false
    }
    // 判断树是否平衡，也就是任意一个节点到叶子节点，经过的黑色链接数量相同
    // 先计算根节点到最左边叶子节点的黑链接数量
    blackNum := 0
    x := tree.Root
    for x != nil {
        if !IsRed(x) { // 是黑色链接
            blackNum = blackNum + 1
        }
        x = x.Left
    }
    if !tree.Root.IsBalanced(blackNum) {
        return false
    }
    return true
}
// 节点所在的子树是否是一棵二分查找树
func (node *LLRBTNode) IsBST() bool {
    if node == nil {
        return true
    }
    // 左子树非空，那么根节点必须大于左儿子节点
    if node.Left != nil {
        if node.Value > node.Left.Value {
        } else {
            fmt.Printf("father:%#v,lchild:%#v,rchild:%#v\n", node, node.Left, node.Right)
            return false
        }
    }
    // 右子树非空，那么根节点必须小于右儿子节点
    if node.Right != nil {
        if node.Value < node.Right.Value {
        } else {
            fmt.Printf("father:%#v,lchild:%#v,rchild:%#v\n", node, node.Left, node.Right)
            return false
        }
    }
    // 左子树也要判断是否是平衡查找树
    if !node.Left.IsBST() {
        return false
    }
    // 右子树也要判断是否是平衡查找树
    if !node.Right.IsBST() {
        return false
    }
    return true
}
// 节点所在的子树是否遵循2-3树
func (node *LLRBTNode) Is23() bool {
    if node == nil {
        return true
    }
    // 不允许右倾红链接
    if IsRed(node.Right) {
        fmt.Printf("father:%#v,rchild:%#v\n", node, node.Right)
        return false
    }
    // 不允许连续两个左红链接
    if IsRed(node) && IsRed(node.Left) {
        fmt.Printf("father:%#v,lchild:%#v\n", node, node.Left)
        return false
    }
    // 左子树也要判断是否遵循2-3树
    if !node.Left.Is23() {
        return false
    }
    // 右子树也要判断是否是遵循2-3树
    if !node.Right.Is23() {
        return false
    }
    return true
}
// 节点所在的子树是否平衡，是否有 blackNum 个黑链接
func (node *LLRBTNode) IsBalanced(blackNum int) bool {
    if node == nil {
        return blackNum == 0
    }
    if !IsRed(node) {
        blackNum = blackNum - 1
    }
    if !node.Left.IsBalanced(blackNum) {
        fmt.Println("node.Left to leaf black link is not ", blackNum)
        return false
    }
    if !node.Right.IsBalanced(blackNum) {
        fmt.Println("node.Right to leaf black link is not ", blackNum)
        return false
    }
    return true
}
```

## 完整程序
`LLRBTree`

# 应用场景
红黑树可以用来作为字典Map的基础数据结构,可以存储键值对,然后通过一个键,可以快速找到键对应的值,相比哈希表查找,不需要占用额外的空间,我们以上的代码实现只有`value`,没有`key:value`,可以简单改造实现字典.