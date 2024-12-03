(def sample-input (slurp "01-sample.txt"))
(def actual-input (slurp "01-actual.txt"))
(defn parse-int [s] (Integer/parseInt s))
(defn diff-lists [input]
  (let [lines (str/split input #"\n")
        pairs (map #(str/split % #"\s") lines)
        left  (sort (map parse-int (map first pairs)))
        right (sort (map parse-int (map last pairs)))]
    (reduce + (map abs (map #(- %1 %2) right left)))))

(println "Part 1:")
(println 11      (diff-lists sample-input))
(println 2113135 (diff-lists actual-input))
(println)

