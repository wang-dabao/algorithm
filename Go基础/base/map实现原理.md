##hash
[深入浅出hashmap](https://segmentfault.com/a/1190000039101378)

* 设计原理
    * 哈希函数：hash函数的性能影响hash表的读写性能，hash函数的输入必将大于输出，所以需要选择一个尽可能均匀分布的hash函数。
    * 冲突解决：拉链法（数组+链表实现，装载因子=元素数量÷桶数量）、开放寻址法（数组实现；其中：装载因子影响hash的读写效率，大于70%，hash表性能急剧下降。装载因子 = 数组中元素的数量 : 数组大小）
    *

* 数据结构
    * 使用了多个数据结构组合表示哈希表，其中 runtime.hmap 是最核心的结构体，
    ```
    // runtime.hmap:
        type hmap struct {
            count     int   //表示当前哈希表中的元素数量；
            B         uint8 //表示当前哈希表持有的 buckets 数量，但是因为哈希表中桶的数量都 2 的倍数，所以该字段会存储对数，也就是 len(buckets) == 2^B；
            hash0     uint32    //hash因子

            buckets    unsafe.Pointer   //桶，即：bmap; 指向一个数组(连续内存空间),数组的类型为[]bmap，这个字段我们可以称之为正常桶。hash键值对存放在桶中
            oldbuckets unsafe.Pointer   //扩容时，存放之前的buckets(Map扩容相关字段)，它的大小是当前 buckets 的一半；
            nevacuate  uintptr          //分流次数，成倍扩容分流操作计数的字段(Map扩容相关字段)

            extra *mapextra     //溢出桶结构，正常桶里面某个bmap存满了，会使用这里面的内存空间存放键值对
            noverflow uint16    //溢出桶里bmap大致的数量
            flags     uint8     //状态标识，比如正在被写、buckets和oldbuckets在被遍历、等量扩容(Map扩容相关字段)

        }

        type mapextra struct {
            overflow    *[]*bmap    //溢出桶, 当正常桶bmap存满了的时候就使用hmap.extra.overflow的bmap;
            oldoverflow *[]*bmap    //扩容时存放之前的overflow(Map扩容相关字段)
            nextOverflow *bmap      //指向溢出桶里下一个可以使用的bmap
        }
        type bmap struct {
            tophash [bucketCnt]uint8    //存储：键的hash的高8位，通过比较不同键的哈希的高 8 位可以减少访问键值对次数以提高性能：
        }

    // cmd/compile/internal/gc.bmap: 编译期间bmap结构体
    type bmap struct {
        topbits  [8]uint8           //长度为8的数组，元素为：key获取的hash的高8位，遍历时对比使用，提高性能。
        keys     [8]keytype         //key数组
        values   [8]valuetype       //val数组
        pad      uintptr            //对齐内存使用的，不是每个bmap都有会这个字段，需要满足一定条件
        overflow uintptr            //指向的hmap.extra.overflow溢出桶里的bmap，上面的字段topbits、keys、elems长度为8，最多存8组键值对，存满了就往指向的这个bmap里存
    }
    ```
    * 字段：
        * count： 表示当前哈希表中的元素数量；
        * flag： 状态标识，比如正在被写、buckets和oldbuckets在被遍历、等量扩容(Map扩容相关字段)
        * B: 表示当前哈希表持有的 buckets 数量，但是因为哈希表中桶的数量都 2 的倍数，所以该字段会存储对数，也就是 len(buckets) == 2^B；
        * hash0： hash因子，它能为哈希函数的结果引入随机性，这个值在创建哈希表时确定，并在调用哈希函数时作为参数传入；
        * buckets: 指向一个数组(连续内存空间),数组的类型为[]bmap，这个字段我们可以称之为正常桶。
        * oldbuckets: 扩容时，存放之前的buckets(Map扩容相关字段)，它的大小是当前 buckets 的一半；
        * extra: 溢出桶结构，正常桶里面某个bmap存满了，会使用这里面的内存空间存放键值对它能够减少扩容的频率所以一直使用至今。
        * noverflow：溢出桶里bmap大致的数量
        * nevacuate：分流次数，成倍扩容分流操作计数的字段(Map扩容相关字段)
        * runtime.bmap: runtime.hmap 的桶类型。每一个 runtime.bmap 都能存储 8 个键值对，当哈希表中存储的数据过多，单个桶已经装满时就会使用 extra.nextOverflow 中桶存储溢出的数据。正常桶和溢出桶在内存中是连续存储中。
        * compile.bmap: 有两个数组分别存放key和value
    * 说明：
        * hmap 和 bmap 的关系：1个 hmap 对应 1个 bmap 数组
        * 每个 bmap 有2个数组，key数组，val数组；
        * 问题: 正常桶hmap.buckets里的bmap是怎么关联上溢出桶hmap.extra.overflow的bmap呢？
        答：就是我们介绍bmap结构时里的bmap.overflow字段(如下图所示)。bmap.overflow是个指针类型，存放了对应使用的溢出桶hmap.extra.overflow里的bmap的地址。
        问：正常桶hmap.buckets里的bmap是什么时候关联上溢出桶hmap.extra.overflow的bmap呢？
        答：Map写操作的时候,看代码：
* 初始化
    * 字面量
        1. 长度 <= 25 时， 代码转换，所有键值对一次性加入hash表中。初始化方式类似于数组和切片；
        2. 长度大于25时，编译器会创建两个数组分别存储键和值，这些键值对会通过如下所示的 for 循环加入哈希。
        ```
        vstatk := []string{"1", "2", "3", ... ， "26"}
        vstatv := []int{1, 2, 3, ... , 26}
        for i := 0; i < len(vstak); i++ {
            hash[vstatk[i]] = vstatv[i]
        }
        ```
    * 运行时
        1. 长度 <= 8 时，使用如下方式快速初始化。
        2. 否则：
            * 计算hash占用的内存是否溢出或者超出能分配的最大值
            * 调用 runtime.fastrand 获取一个随机的哈希种子
            * 根据传入的 hint 计算需要的桶的数量；
            * 使用 runtime.makeBucketArray 创建用于保存桶的数组。（根据传入的 B 计算出的需要创建的桶数量并在内存中分配一片连续的空间用于存储数据）
                * 当桶的数量小于 24 时，由于数据较少、使用溢出桶的可能性较低，会省略创建的过程以减少额外开销；
                * 当桶的数量多于 24 时，会额外创建 2𝐵−4 个溢出桶；（正常桶和溢出桶在内存中的存储空间是连续的，只是被 runtime.hmap 中的不同字段引用，当溢出桶数量较多时会通过 runtime.newobject 创建新的溢出桶）


* 读写操作
    * 访问：
    ```
    //方法1： 通过下标访问
      _ = hash[key]

    //方法2： 依次遍历
      for k, v := range hash {
          // k, v
      }
    ```

    * 下标访问：需要知道单个键值key. 通过计算key的hash值，然后与桶的数据取模，确定桶的位置，然后遍历单链表，通过比对key值，获取val
        1. 在编译的类型检查期间，hash[key] 以及类似的操作都会被转换成哈希的 OINDEXMAP操作。根据左侧参数个数决定使用的运行时方法。
        2. runtime.mapaccess1：
           * 通过hash函数、哈希因子计算当前key的hash值，
           * 哈希值和buckets数组的长度通过位操作获取数组位置的索引，
                通过 runtime.bucketMask() , runtime.add 拿到该键值对所在的桶序号和哈希高位的 8 位数字。
        3. 遍历bmap里的键，和目标key对比获取key的索引，根据key的索引通过计算偏移量，获取到对应value
            * 在 bucketloop 循环中，先遍历正常桶，依次比较桶中的 tophash 值和key的高8位hash值是否相等，找到高8位hash值时，就拿到了索引位置i的key, 然后判断该位置的key是否和目标key相等，根据索引位置获取值。
            * 如果没有找到值是，然后判断当前bmap的overflow值是否为nil, 不为nil时，重复步骤1继续遍历溢出桶中的数据，
              用于选择桶序号的是哈希的最低几位，而用于加速访问的是哈希的高 8 位，这种设计能够减少同一个桶中有大量相等 tophash 的概率影响性能。

            * 扩容期间获取键值对的逻辑：当哈希表的 oldbuckets 存在时，会先定位到旧桶并在该桶没有被分流时从中获取键值对。

           * runtime.mapaccess2: 只是在 runtime.mapaccess1 的基础上多返回了一个标识键值对是否存在的 bool 值：
        4. 案例：key的查找过程：
            [案例图片](https://segmentfault.com/img/remote/1460000023879183)
            假定 B = 5，所以 bucket 总数就是 2^5 = 32。首先计算出待查找 key 的哈希，使用低 5 位 00110，找到对应的 6 号 bucket，使用高 8 位 10010111，对应十进制 151，在 6 号 bucket 中寻找 tophash 值（HOB hash）为 151 的 key，找到了 2 号槽位，这样整个查找过程就结束了。
            如果在 bucket 中没找到，并且 overflow 不为空，还要继续去 overflow bucket 中寻找，直到找到或是所有的 key 槽位都找遍了，包括所有的 overflow bucket。
    * 遍历：可以获取全部键值对。 通过range方式获取。
    * 写入（runtime.mapassign）
        * 1. 根据传入的键拿到对应的哈希值和桶；
        * 2. 遍历比较：桶中存储的 tophash 和键的 hash，如果找到了相同结果就会返回目标位置的地址。其中 inserti 表示目标元素的在桶中的索引，insertk 和 val 分别表示键值对的地址，获得目标地址之后会通过算术计算寻址获得键值对 k 和 val：
             会依次遍历正常桶和溢出桶中存储的数据，整个过程会分别判断 tophash 是否相等、key 是否相等，遍历结束后会从循环中跳出。
        * 3. 如果当前桶满了，创建新的溢出桶（runtime.hmap.newoverflow）或者在已有的溢出桶中保存数据。（新创建的同会追加到已有桶的默认，还会增加hash表中noverflow计数器）
        * 4. 如果当前键值对在hash表中不存在，会为新键值对规划存储的内存地址。通过：runtime.typedmemmove 将键移动到对应的内存空间并返回对应值的地址cval.
        * 5. 如果键值对在hash中存在，直接返回目标区域的内存地址，并不会在 runtime.mapassign 中将值拷贝到桶，在编译期间操作赋值。

        * 特殊说明：当哈希表正在处于扩容状态时，每次向哈希表写入值时都会触发 runtime.growWork 增量拷贝哈希表中的内容：

    * 扩容（在扩容的过程中要保证哈希的访问是比较有意思的话题）
        * 总结：哈希在存储元素过多时会触发扩容操作，每次都会将桶的数量翻倍，扩容过程不是原子的，而是通过 runtime.growWork 增量触发的，在扩容期间访问哈希表时会使用旧桶，向哈希表写入数据时会触发旧桶元素的分流。
        除了这种正常的扩容之外，为了解决大量写入、删除造成的内存泄漏问题，哈希引入了 sameSizeGrow 这一机制，在出现较多溢出桶时会整理哈希的内存减少空间的占用。
        * 特点：不是原子过程，runtime.mapassign 还需要判断当前hash是否处于扩容状态，避免二次扩容造成混乱。
        * runtime.mapassign `函数触发扩容的条件`：
            * 装载因子已经超过 6.5；(装载因子=元素数量:桶数量) ===> *翻倍扩容*
                装载因子过大表示哈希空间利用率过高，容易引发哈希冲突。查找效率和插入效率都变低了
            * overflow 的 桶数量过多, 哈希使用了太多溢出桶；==> *等量扩容* sameSizeGrow
                B<=15，已使用的溢出桶个数>=2的B次方时，引发等量扩容。
                B>15，已使用的溢出桶个数>=2的15次方时，引发等量扩容。
                当我们持续向哈希中插入数据并将它们全部删除时，如果哈希表中的数据量没有超过阈值，就会不断积累溢出桶造成缓慢的内存泄漏。
                overflow bucket 数量太多，导致 key 会很分散，查找插入效率低得吓人

        * 等量扩容：
            * 场景：一旦哈希中出现了过多的溢出桶，它会创建新桶保存数据，垃圾回收会清理老的溢出桶并释放内存。
            * 特点：创建出的新桶数量和旧桶一样，该函数只是创建了新的桶，并没有对数据进行拷贝和转移。

        * 扩容过程: runtime.hashGrow代码如下
            * 1. 通过 runtime.makeBucketArray 创建一组新桶和预创建的溢出桶。桶的数量是原来桶数量的2倍。B=B+1
            * 2. 随后将原有的桶数组指向 oldbuckets 上并将新的空桶设置到 buckets 上
            * 3. 溢出桶也使用了相同的逻辑更新
            ```
            func hashGrow(t *maptype, h *hmap) {
            	bigger := uint8(1)
            	if !overLoadFactor(h.count+1, h.B) {    //溢出桶过多，触发等量扩容；
            		bigger = 0
            		h.flags |= sameSizeGrow
            	}
            	oldbuckets := h.buckets
            	newbuckets, nextOverflow := makeBucketArray(t, h.B+bigger, nil)

            	h.B += bigger
            	h.flags = flags
            	h.oldbuckets = oldbuckets
            	h.buckets = newbuckets
            	h.nevacuate = 0
            	h.noverflow = 0

            	h.extra.oldoverflow = h.extra.overflow
            	h.extra.overflow = nil
            	h.extra.nextOverflow = nextOverflow
            }
            ```
        * 数据分流过程：runtime.evacuate
            * 翻量扩容：runtime.evacuate 会将一个旧桶中的数据分流到两个新桶，所以它会创建两个用于保存分配上下文的 runtime.evacDst 结构体，这两个结构体分别指向了一个新桶
            * 等量扩容：创建出的新桶数量和旧桶一样，所以两个 runtime.evacDst 只会初始化一个
            * 过程：
                * 调用 runtime.advanceEvacuationMark 增加哈希的 nevacuate 计数器，并在所有的旧桶都被分流后清空哈希的 oldbuckets 和 oldoverflow：
        说明：
            触发条件2：溢出桶过多时，其实元素没那么多，但是 overflow bucket 数特别多，说明很多 bucket 都没装满。
            解决办法就是开辟一个新 bucket 空间，将老 bucket 中的元素移动到新 bucket，使得同一个 bucket 中的 key 排列地更紧密。
            这样，原来，在 overflow bucket 中的 key 可以移动到 bucket 中来。结果是节省空间，提高 bucket 利用率，map 的查找和插入效率自然就会提升。

            * 对于触发条件1：装载因子过大时，需要增量扩容，扩容后数据搬迁时，需要重新计算key的哈希，才能决定他落在那个桶。
            例如，原来 B = 5，计算出 key 的哈希后，只用看它的低 5 位，就能决定它落在哪个 bucket。扩容后，B 变成了 6，因此需要多看一位，它的低 6 位决定 key 落在哪个 bucket。这称为 rehash。

            * 理解了上面 bucket 序号的变化，我们就可以回答另一个问题了：为什么遍历 map 是无序的？
            map 在扩容后，会发生 key 的搬迁，原来落在同一个 bucket 中的 key，搬迁后，有些 key 就要远走高飞了（bucket 序号加上了 2^B）。而遍历的过程，就是按顺序遍历 bucket，同时按顺序遍历 bucket 中的 key。搬迁后，key 的位置发生了重大的变化，有些 key 飞上高枝，有些 key 则原地不动。这样，遍历 map 的结果就不可能按原来的顺序了。
            多说一句，“迭代 map 的结果是无序的”这个特性是从 go 1.0 开始加入的。每次都是从一个随机值序号的 bucket 开始遍历，并且是从这个 bucket 的一个随机序号的 cell 开始遍历。这样，即使你是一个写死的 map，仅仅只是遍历它，也不太可能会返回一个固定序列的 key/value 对了。

            * 再明确一个问题：如果扩容后，B 增加了 1，意味着 buckets 总数是原来的 2 倍，原来 1 号的桶“裂变”到两个桶。

            * 渐进式搬迁：每次只搬迁1个桶，每个桶的搬迁分为正常桶搬迁+溢出桶搬迁。
            搬迁是一个“渐进”过程，并不会一下子就全部搬迁完毕。所以在搬迁过程中，oldbuckets 指针还会指向原来老的 []bmap，并且已经搬迁完毕的 key 的 tophash 值会是一个状态值，表示 key 的搬迁去向。
            evacuate 函数每次只完成一个 bucket 的搬迁工作，因此要遍历完此 bucket 的所有的 cell，将有值的 cell copy 到新的地方。bucket 还会链接 overflow bucket，它们同样需要搬迁。
            因此会有 2 层循环，外层遍历 bucket 和 overflow bucket，内层遍历 bucket 的所有 cell。这样的循环在 map 的源码里到处都是，要理解透了。

            *特殊key,特殊处理：有一种 key，每次对它计算 hash，得到的结果都不一样。这个 key 就是 math.NaN() 的结果，它的含义是 not a number，类型是 float64。当它作为 map 的 key，在搬迁的时候，会遇到一个问题：再次计算它的哈希值和它当初插入 map 时的计算出来的哈希值不一样！
             你可能想到了，这样带来的一个后果是，这个 key 是永远不会被 Get 操作获取的！当我使用 m[math.NaN()] 语句的时候，是查不出来结果的。这个 key 只有在遍历整个 map 的时候，才有机会现身。所以，可以向一个 map 插入任意数量的 math.NaN() 作为 key。
             当搬迁碰到 math.NaN() 的 key 时，只通过 tophash 的最低位决定分配到 X part 还是 Y part（如果扩容后是原来 buckets 数量的 2 倍）。如果 tophash 的最低位是 0 ，分配到 X part；如果是 1 ，则分配到 Y part。

    * 删除（delete(hash, key)）
        * 编译期间，delete 关键字会被转换为 ODELETE 节点，而 cmd/compile/internal/gc.walkexpr 会将 ODELETE 节点转换成 runtime.mapdelete 函数簇中的一个
        * 删除逻辑与写入逻辑很相似
        * 如果在删除期间遇到了哈希表的扩容，就会分流桶中的元素，分流结束之后会找到桶中的目标元素完成键值对的删除工作。















