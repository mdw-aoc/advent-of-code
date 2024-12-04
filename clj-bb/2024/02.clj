(defn parse-int [s] (Integer/parseInt s))
(defn prep-line [line]
  (map parse-int (str/split line #"\s")))
(defn prep-input [filename]
  (as-> filename $
        (slurp $)
        (str/split $ #"\n")
        (map prep-line $)))
(def sample-input (prep-input "02-sample.txt"))
(def actual-input (prep-input "02-actual.txt"))

(defn in-range? [n] (and (> n -4) (< n 4)))
(defn is-safe? [ints]
  (let [diffs (map #(- (first %1) (second %1)) (partition 2 1 ints))]
    (and (or (every? pos? diffs) (every? neg? diffs))
         (every? in-range? diffs))))
(defn part1 [input] (->> input (filter is-safe?) count))
(println 2   (part1 sample-input))
(println 526 (part1 actual-input))

(defn without [coll index]
  (concat (take index coll) (drop (inc index) coll)))
(defn has-safe-permutations? [ints]
  (->> (range (count ints))
       (map #(without ints %))
       (filter is-safe?) count pos?))
(defn part2 [input] (->> input (filter has-safe-permutations?) count))
(println 4   (part2 sample-input))
(println 566 (part2 actual-input))