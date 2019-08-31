
DROP TABLE IF EXISTS `user`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `user` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `userName` varchar(16) NOT NULL,
  `password` varchar(32) NOT NULL,
  `name` varchar(16) DEFAULT NULL,
  `age` int(3) DEFAULT NULL,
  `phone` varchar(16) DEFAULT NULL,
  `date` date NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=8 DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `user`
--

LOCK TABLES `user` WRITE;
/*!40000 ALTER TABLE `user` DISABLE KEYS */;
INSERT INTO `user` VALUES (1,'admin','123456','admin',27,'13634450823','2018-04-04'),(2,'fyx912','123456','fyx912',20,'13532550872','2018-04-08'),(3,'ding','123456','ding',21,'13532550872','2018-04-08'),(4,'tintin','123456','Tintin',21,'13532550872','2018-04-08'),(5,'Jon','654321','乔恩',24,'14354329861','2018-04-08'),(6,'james','123456','james',31,'13532550872','2018-04-08'),(7,'fixe','123456','菲克斯',31,'13532550872','2018-04-08');
/*!40000 ALTER TABLE `user` ENABLE KEYS */;
UNLOCK TABLES;