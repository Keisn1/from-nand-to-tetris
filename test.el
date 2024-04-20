
(org-link-decode "e6ac5e60-8fc9-44d0-a457-3fc841345c10")

(org-link-decode "Encoded%20String")

(org-roam-ui--get-text 'e6ac5e60-8fc9-44d0-a457-3fc841345c10)
(setq id "e6ac5e60-8fc9-44d0-a457-3fc841345c10")
(insert (format "The value of my-variable is: %s" id))
The value of my-variable is: e6ac5e60-8fc9-44d0-a457-3fc841345c10


;; set node
(setq node (org-roam-populate (org-roam-node-create :id id)))
(insert (format "The value of my-variable is: %s" node))
;; The value of my-variable is:
;; #s(org-roam-node /home/kay/workspace/from-nand-to-tetris/org-roam-files/20240415121719-m_register.org M register nil
;;                  (26147 48361 927989 867000)
;;                  (26141 6895 251013 405000)
;;                  e6ac5e60-8fc9-44d0-a457-3fc841345c10
;;                  0 1 nil nil nil nil M register
;;                  ((CATEGORY . 20240415121719-m_register)
;;                   (ID . e6ac5e60-8fc9-44d0-a457-3fc841345c10)
;;                   (BLOCKED . )
;;                   (FILE . /home/kay/workspace/from-nand-to-tetris/org-roam-files/20240415121719-m_register.org)
;;                   (PRIORITY . B)) nil nil nil nil)

;; set file
(setq file (org-roam-node-file node))
(insert (format "The value of my-variable is: %s" file))
;; The value of file is: /home/kay/workspace/from-nand-to-tetris/org-roam-files/20240415121719-m_register.org

(org-roam-with-temp-buffer
    file
  (when (> (org-roam-node-level node) 0)
    ;; Heading nodes have level 1 and greater.
    (goto-char (org-roam-node-point node))
    (org-narrow-to-element))
  (buffer-substring-no-properties (buffer-end -1) (buffer-end 1)))



(org-roam-ui--get-text "e6ac5e60-8fc9-44d0-a457-3fc841345c10")

(org-roam-ui--get-text (org-link-decode "e6ac5e60-8fc9-44d0-a457-3fc841345c10"))

(org-roam-ui--get-text "e6ac5e60-8fc9-44d0-a457-3fc841345c10")
(setq id "e6ac5e60-8fc9-44d0-a457-3fc841345c10")


(insert (format "%s" (org-roam-node-create :id id)))
#s(org-roam-node nil nil nil nil nil e6ac5e60-8fc9-44d0-a457-3fc841345c10 nil nil nil nil nil nil nil nil nil nil nil nil)

(insert (format "%s" (org-roam-populate (org-roam-node-create :id id))))
#s(org-roam-node /home/kay/workspace/from-nand-to-tetris/org-roam-files/20240415121719-m_register.org M register nil (26147 48361 927989 867000) (26141 6895 251013 405000) e6ac5e60-8fc9-44d0-a457-3fc841345c10 0 1 nil nil nil nil M register ((CATEGORY . 20240415121719-m_register) (ID . e6ac5e60-8fc9-44d0-a457-3fc841345c10) (BLOCKED . ) (FILE . /home/kay/workspace/from-nand-to-tetris/org-roam-files/20240415121719-m_register.org) (PRIORITY . B)) nil nil nil nil)

(setq node (org-roam-populate (org-roam-node-create :id id)))
(insert (format "The value of my-variable is: %s" node))
The value of my-variable is:
#s(org-roam-node /home/kay/workspace/from-nand-to-tetris/org-roam-files/20240415121719-m_register.org M register nil (26147 48361 927989 867000) (26141 6895 251013 405000) e6ac5e60-8fc9-44d0-a457-3fc841345c10
                 0 1 nil nil nil nil M register ((CATEGORY . 20240415121719-m_register) (ID . e6ac5e60-8fc9-44d0-a457-3fc841345c10) (BLOCKED . ) (FILE . /home/kay/workspace/from-nand-to-tetris/org-roam-files/20240415121719-m_register.org) (PRIORITY . B)) nil nil nil nil)

(setq file (org-roam-node-file node))

(insert (format "The value of my-variable is: %s" file))
The value of my-variable is:
/home/kay/workspace/from-nand-to-tetris/org-roam-files/20240415121719-m_register.org
