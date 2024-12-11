(defn prep-input [filename] (str/split (slurp filename) #"\n"))
(def sample-input (prep-input "04-sample.txt"))
(def actual-input (prep-input "04-actual.txt"))

(defn resolve-char [rows start offset]
  (let [row (nth rows (+ (first start) (first offset)) nil)
        col (nth row (+ (second start) (second offset)) nil)]
    col))

(defn slice [rows coords start]
  (let [chars (map (partial resolve-char rows start) coords)]
    (when (every? (complement nil?) chars)
      (apply str chars))))

(defn all-coords [rows]
  (for [row (range (count rows))
        col (range (count (first rows)))]
    [row col]))

(defn pattern-matches [pattern rows]
  (->> rows
       all-coords
       (map (partial slice rows (:coords pattern)))
       (remove nil?)
       (filter #(= (:expect pattern) %))))

(defn rotate90 [rows]
  (reverse
    (for [column (range (count (first rows)))]
      (apply str (map #(nth % column) rows)))))

(defn rotated-pattern-matches [rows pattern]
  (->> rows
       (iterate rotate90)
       (take 4)
       (map (partial pattern-matches pattern))))

(def part1-horizontal
  {:coords [[0 0] [0 1] [0 2] [0 3]]
   :expect "XMAS"})

(def part1-diagonal
  {:coords [[0 0] [1 1] [2 2] [3 3]]
   :expect "XMAS"})

(defn all-matches [rows & patterns]
  (->> patterns
       (map (partial rotated-pattern-matches rows))
       flatten
       count))

(println 18   (all-matches sample-input part1-horizontal part1-diagonal))
(println 2545 (all-matches actual-input part1-horizontal part1-diagonal))

(def part2-cross
  {:coords [[0 0] ,,,,, [0 2]
                  [1 1]
            [2 0] ,,,,, [2 2]]
   :expect "MMASS"})

(println 9    (all-matches sample-input part2-cross))
(println 1886 (all-matches actual-input part2-cross))

