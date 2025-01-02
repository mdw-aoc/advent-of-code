(def sample-input (slurp "08-sample.txt"))
(def actual-input (slurp "08-actual.txt"))

(defn world-set [input]
  (set
    (let [lines (str/split input #"\n")]
       (for [row (range (count lines))
             col (range (count (first lines)))]
         [row col]))))

(defn locate-antennas [input]
  (loop [input input row 0 col 0 result {}]
    (if (empty? input)
      result
      (let [char   (first input)
            input  (rest input)
            point  [row col]
            row    (if (= char \newline) (inc row) row)
            col    (if (= char \newline) 0 (inc col))
            result (if (or (= char \.) (= char \newline))
                     result
                     (update result char conj point))]
        (recur input row col result)))))

(defn locate-antinodes [antennas]
  (for [a antennas
        b antennas
        :when (not (= a b))
        :let [d1 (mapv - a b)]]
    (mapv + d1 a)))

(defn part1 [input]
  (->> input
       locate-antennas seq (map last)
       (mapcat locate-antinodes) set
       (filter (world-set input))
       sort count))

(println 14  (part1 sample-input))
(println 311 (part1 actual-input))
