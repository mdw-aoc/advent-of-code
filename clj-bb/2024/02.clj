(def sample-input (slurp "02-sample.txt"))
(def actual-input (slurp "02-actual.txt"))
(defn parse-int [s] (Integer/parseInt s))
(defn in-range? [n] (and (> n -4) (< n 4)))
(defn is-safe? [line]
  (let [ints  (map parse-int (str/split line #"\s"))
        diffs (map #(- (first %1) (second %1)) (partition 2 1 ints))]
    (and (or (every? pos? diffs) (every? neg? diffs))
         (every? in-range? diffs))))
(defn part1 [input]
  (count (filter is-safe? (str/split input #"\n"))))
(println 2   (part1 sample-input))
(println 526 (part1 actual-input))
