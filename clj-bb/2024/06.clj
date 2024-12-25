(def right-turns
  {[-1 0] [0 1]
   [0 1]  [1 0]
   [1 0]  [0 -1]
   [0 -1] [-1 0]})

(defn new-state []
  {:obstacles   (set [])
   :path        []
   :start       [-1 -1]
   :at          [-1 -1]
   :facing      [-1  0]
   :lower-right [-1 -1]
   :visited     (set []) ; obstacles visited (for detecting loops))
   :looped?     false}) 

(defn parse-input [input]
  (let [state (new-state)]
    (loop [state state
           input input
           row 0 col 0]
      (let [char (first input)
            input (rest input)]
        (if (nil? char)
          (assoc state :lower-right [row (dec col)])
          (case char
            \newline (recur state input (inc row) 0)
            \.       (recur state input row (inc col))
            \^       (recur (-> state
                                 (assoc :start [row col] :at [row col])
                                 (update-in [:path] conj [row col]))
                            input row (inc col))
            \#       (recur (-> state
                                (update-in [:obstacles] conj [row col]))
                            input row (inc col))))))))

(defn move [from direction]
  (vec (map + from direction)))

(defn patrol [state]
  (let [at        (:at state) 
        facing    (:facing state)
        obstacles (:obstacles state)
        path      (:path state)
        visited   (:visited state)
        upcoming  (move at facing)
        path      (conj path at)]
    (if (not (contains? obstacles upcoming))
      (assoc state :at upcoming :path path)
      (let [now-facing (get right-turns facing)
            looped?    (or (:looped? state) (contains? visited [at upcoming]))
            visited    (conj visited [at upcoming])]
        (assoc state :facing now-facing
                     :path path
                     :visited visited
                     :looped? looped?)))))

(defn in-bounds? [{:keys [at lower-right]}]
  (let [[at-row at-col] at
        [max-row max-col] lower-right]
    (and (>= at-row 0)
         (>= at-col 0)
         (<= at-row max-row)
         (<= at-col max-col))))

(defn make-rounds [state]
  (->> state
       (iterate patrol)
       (drop-while #(and (in-bounds? %1) (not (:looped? %1))))
       first))

(defn part1 [input]
  (->> input parse-input make-rounds :path set count))

(def sample-input (slurp "06-sample.txt"))
(def actual-input (slurp "06-actual.txt"))

(println 41   (part1 sample-input))
(println 4752 (part1 actual-input))

(defn part2 [input]
  (->> input
       parse-input
       make-rounds
       :path
       set
       (remove #(= % (:start (parse-input input)))) ; remove starting coord
       (map #(update-in (parse-input input) [:obstacles] conj %))
       (map make-rounds)
       (filter :looped?)
       count))
 
(println 6    (part2 sample-input))
(println "Stand by, part 2 takes a while...")
(println 1719 (part2 actual-input)) ; long-running!

