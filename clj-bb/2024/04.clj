(defn prep-input [filename]
  (as-> filename $ (slurp $) (str/split $ #"\n") (str/join "|" $)))
(def sample-input (prep-input "04-sample.txt"))
(def actual-input (prep-input "04-actual.txt"))

(defn upper-diagonal [lines starting-row]
  (let [rows (range starting-row -1 -1)
        cols (range (count rows))]
    (for [[row col] (partition 2 (interleave rows cols))]
      (nth (nth lines row) col))))
(defn upper-diagonals [lines]
  (->> lines count range (map #(upper-diagonal lines %))))
(defn lower-diagonal [lines col]
  (let [cols (range col (count (first lines)))
        rows (range (dec (count lines)) -1 -1)
        pairs (partition 2 (interleave rows cols))]
    (for [[row col] (partition 2 (interleave rows cols))]
      (nth (nth lines row) col))))
(defn lower-diagonals [lines]
  (->> lines first count (range 1) (map #(lower-diagonal lines %))))
(defn rotate45 [text]
  (let [lines (str/split text #"\|")
        upper (map #(apply str %) (upper-diagonals lines))
        lower (map #(apply str %) (lower-diagonals lines))]
    (str/join "|" (concat upper lower))))

(defn rotate90 [text]
  (str/join "|"
    (let [rows (str/split text #"\|")]
      (reverse
        (for [column (range (count (first rows)))]
          (apply str (map #(nth % column) rows)))))))

(defn rotations [text]
  (let [nineties   (take 4 (iterate rotate90 text))
        fortyfives (map rotate45 nineties)]
    (str/join "||" (interleave nineties fortyfives))))

(defn part1 [text]
  (count (re-seq #"XMAS" (rotations text))))

(println 18   (part1 sample-input))
(println 2545 (part1 actual-input))
