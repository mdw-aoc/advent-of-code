(def sample-input (slurp "01-sample.txt"))
(def actual-input (slurp "01-actual.txt"))
(defn parse-int [s] (Integer/parseInt s))
(defn sorted-ints [values] (sort (map parse-int values)))
(defn parse-lists [input]
  (let [lines (str/split input #"\n")
        pairs (map #(str/split % #"\s") lines)]
    {:left  (sorted-ints (map first pairs))
     :right (sorted-ints (map last pairs))}))

(defn diff-lists [input]
  (let [{left :left right :right} (parse-lists input)]
    (reduce + (map abs (map #(- %1 %2) left right)))))

(println "Part 1:")
(println 11      (diff-lists sample-input))
(println 2113135 (diff-lists actual-input))
(println)

(defn score-similarity [input]
  (let [{left :left right :right} (parse-lists input)
        right (frequencies right)]
    (reduce + (map #(* %1 (get right %1 0)) left))))

(println "Part 2:")
(println 31       (score-similarity sample-input))
(println 19097157 (score-similarity actual-input))
(println)